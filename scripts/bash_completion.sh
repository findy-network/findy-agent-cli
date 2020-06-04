# bash completion for findy-cli                            -*- shell-script -*-

__findy-cli_debug()
{
    if [[ -n ${BASH_COMP_DEBUG_FILE} ]]; then
        echo "$*" >> "${BASH_COMP_DEBUG_FILE}"
    fi
}

# Homebrew on Macs have version 1.3 of bash-completion which doesn't include
# _init_completion. This is a very minimal version of that function.
__findy-cli_init_completion()
{
    COMPREPLY=()
    _get_comp_words_by_ref "$@" cur prev words cword
}

__findy-cli_index_of_word()
{
    local w word=$1
    shift
    index=0
    for w in "$@"; do
        [[ $w = "$word" ]] && return
        index=$((index+1))
    done
    index=-1
}

__findy-cli_contains_word()
{
    local w word=$1; shift
    for w in "$@"; do
        [[ $w = "$word" ]] && return
    done
    return 1
}

__findy-cli_handle_reply()
{
    __findy-cli_debug "${FUNCNAME[0]}"
    local comp
    case $cur in
        -*)
            if [[ $(type -t compopt) = "builtin" ]]; then
                compopt -o nospace
            fi
            local allflags
            if [ ${#must_have_one_flag[@]} -ne 0 ]; then
                allflags=("${must_have_one_flag[@]}")
            else
                allflags=("${flags[*]} ${two_word_flags[*]}")
            fi
            while IFS='' read -r comp; do
                COMPREPLY+=("$comp")
            done < <(compgen -W "${allflags[*]}" -- "$cur")
            if [[ $(type -t compopt) = "builtin" ]]; then
                [[ "${COMPREPLY[0]}" == *= ]] || compopt +o nospace
            fi

            # complete after --flag=abc
            if [[ $cur == *=* ]]; then
                if [[ $(type -t compopt) = "builtin" ]]; then
                    compopt +o nospace
                fi

                local index flag
                flag="${cur%=*}"
                __findy-cli_index_of_word "${flag}" "${flags_with_completion[@]}"
                COMPREPLY=()
                if [[ ${index} -ge 0 ]]; then
                    PREFIX=""
                    cur="${cur#*=}"
                    ${flags_completion[${index}]}
                    if [ -n "${ZSH_VERSION}" ]; then
                        # zsh completion needs --flag= prefix
                        eval "COMPREPLY=( \"\${COMPREPLY[@]/#/${flag}=}\" )"
                    fi
                fi
            fi
            return 0;
            ;;
    esac

    # check if we are handling a flag with special work handling
    local index
    __findy-cli_index_of_word "${prev}" "${flags_with_completion[@]}"
    if [[ ${index} -ge 0 ]]; then
        ${flags_completion[${index}]}
        return
    fi

    # we are parsing a flag and don't have a special handler, no completion
    if [[ ${cur} != "${words[cword]}" ]]; then
        return
    fi

    local completions
    completions=("${commands[@]}")
    if [[ ${#must_have_one_noun[@]} -ne 0 ]]; then
        completions=("${must_have_one_noun[@]}")
    fi
    if [[ ${#must_have_one_flag[@]} -ne 0 ]]; then
        completions+=("${must_have_one_flag[@]}")
    fi
    while IFS='' read -r comp; do
        COMPREPLY+=("$comp")
    done < <(compgen -W "${completions[*]}" -- "$cur")

    if [[ ${#COMPREPLY[@]} -eq 0 && ${#noun_aliases[@]} -gt 0 && ${#must_have_one_noun[@]} -ne 0 ]]; then
        while IFS='' read -r comp; do
            COMPREPLY+=("$comp")
        done < <(compgen -W "${noun_aliases[*]}" -- "$cur")
    fi

    if [[ ${#COMPREPLY[@]} -eq 0 ]]; then
		if declare -F __findy-cli_custom_func >/dev/null; then
			# try command name qualified custom func
			__findy-cli_custom_func
		else
			# otherwise fall back to unqualified for compatibility
			declare -F __custom_func >/dev/null && __custom_func
		fi
    fi

    # available in bash-completion >= 2, not always present on macOS
    if declare -F __ltrim_colon_completions >/dev/null; then
        __ltrim_colon_completions "$cur"
    fi

    # If there is only 1 completion and it is a flag with an = it will be completed
    # but we don't want a space after the =
    if [[ "${#COMPREPLY[@]}" -eq "1" ]] && [[ $(type -t compopt) = "builtin" ]] && [[ "${COMPREPLY[0]}" == --*= ]]; then
       compopt -o nospace
    fi
}

# The arguments should be in the form "ext1|ext2|extn"
__findy-cli_handle_filename_extension_flag()
{
    local ext="$1"
    _filedir "@(${ext})"
}

__findy-cli_handle_subdirs_in_dir_flag()
{
    local dir="$1"
    pushd "${dir}" >/dev/null 2>&1 && _filedir -d && popd >/dev/null 2>&1 || return
}

__findy-cli_handle_flag()
{
    __findy-cli_debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"

    # if a command required a flag, and we found it, unset must_have_one_flag()
    local flagname=${words[c]}
    local flagvalue
    # if the word contained an =
    if [[ ${words[c]} == *"="* ]]; then
        flagvalue=${flagname#*=} # take in as flagvalue after the =
        flagname=${flagname%=*} # strip everything after the =
        flagname="${flagname}=" # but put the = back
    fi
    __findy-cli_debug "${FUNCNAME[0]}: looking for ${flagname}"
    if __findy-cli_contains_word "${flagname}" "${must_have_one_flag[@]}"; then
        must_have_one_flag=()
    fi

    # if you set a flag which only applies to this command, don't show subcommands
    if __findy-cli_contains_word "${flagname}" "${local_nonpersistent_flags[@]}"; then
      commands=()
    fi

    # keep flag value with flagname as flaghash
    # flaghash variable is an associative array which is only supported in bash > 3.
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        if [ -n "${flagvalue}" ] ; then
            flaghash[${flagname}]=${flagvalue}
        elif [ -n "${words[ $((c+1)) ]}" ] ; then
            flaghash[${flagname}]=${words[ $((c+1)) ]}
        else
            flaghash[${flagname}]="true" # pad "true" for bool flag
        fi
    fi

    # skip the argument to a two word flag
    if [[ ${words[c]} != *"="* ]] && __findy-cli_contains_word "${words[c]}" "${two_word_flags[@]}"; then
			  __findy-cli_debug "${FUNCNAME[0]}: found a flag ${words[c]}, skip the next argument"
        c=$((c+1))
        # if we are looking for a flags value, don't show commands
        if [[ $c -eq $cword ]]; then
            commands=()
        fi
    fi

    c=$((c+1))

}

__findy-cli_handle_noun()
{
    __findy-cli_debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"

    if __findy-cli_contains_word "${words[c]}" "${must_have_one_noun[@]}"; then
        must_have_one_noun=()
    elif __findy-cli_contains_word "${words[c]}" "${noun_aliases[@]}"; then
        must_have_one_noun=()
    fi

    nouns+=("${words[c]}")
    c=$((c+1))
}

__findy-cli_handle_command()
{
    __findy-cli_debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"

    local next_command
    if [[ -n ${last_command} ]]; then
        next_command="_${last_command}_${words[c]//:/__}"
    else
        if [[ $c -eq 0 ]]; then
            next_command="_findy-cli_root_command"
        else
            next_command="_${words[c]//:/__}"
        fi
    fi
    c=$((c+1))
    __findy-cli_debug "${FUNCNAME[0]}: looking for ${next_command}"
    declare -F "$next_command" >/dev/null && $next_command
}

__findy-cli_handle_word()
{
    if [[ $c -ge $cword ]]; then
        __findy-cli_handle_reply
        return
    fi
    __findy-cli_debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"
    if [[ "${words[c]}" == -* ]]; then
        __findy-cli_handle_flag
    elif __findy-cli_contains_word "${words[c]}" "${commands[@]}"; then
        __findy-cli_handle_command
    elif [[ $c -eq 0 ]]; then
        __findy-cli_handle_command
    elif __findy-cli_contains_word "${words[c]}" "${command_aliases[@]}"; then
        # aliashash variable is an associative array which is only supported in bash > 3.
        if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
            words[c]=${aliashash[${words[c]}]}
            __findy-cli_handle_command
        else
            __findy-cli_handle_noun
        fi
    else
        __findy-cli_handle_noun
    fi
    __findy-cli_handle_word
}

_findy-cli_agency_ping()
{
    last_command="findy-cli_agency_ping"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--base-addr=")
    two_word_flags+=("--base-addr")
    local_nonpersistent_flags+=("--base-addr=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_agency_start()
{
    last_command="findy-cli_agency_start"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--a2a=")
    two_word_flags+=("--a2a")
    local_nonpersistent_flags+=("--a2a=")
    flags+=("--did=")
    two_word_flags+=("--did")
    local_nonpersistent_flags+=("--did=")
    flags+=("--genesis=")
    two_word_flags+=("--genesis")
    local_nonpersistent_flags+=("--genesis=")
    flags+=("--hostaddr=")
    two_word_flags+=("--hostaddr")
    local_nonpersistent_flags+=("--hostaddr=")
    flags+=("--hostport=")
    two_word_flags+=("--hostport")
    local_nonpersistent_flags+=("--hostport=")
    flags+=("--pool=")
    two_word_flags+=("--pool")
    local_nonpersistent_flags+=("--pool=")
    flags+=("--protocol=")
    two_word_flags+=("--protocol")
    local_nonpersistent_flags+=("--protocol=")
    flags+=("--psmdb=")
    two_word_flags+=("--psmdb")
    local_nonpersistent_flags+=("--psmdb=")
    flags+=("--register=")
    two_word_flags+=("--register")
    local_nonpersistent_flags+=("--register=")
    flags+=("--reset")
    local_nonpersistent_flags+=("--reset")
    flags+=("--seed=")
    two_word_flags+=("--seed")
    local_nonpersistent_flags+=("--seed=")
    flags+=("--serverport=")
    two_word_flags+=("--serverport")
    local_nonpersistent_flags+=("--serverport=")
    flags+=("--service-name=")
    two_word_flags+=("--service-name")
    local_nonpersistent_flags+=("--service-name=")
    flags+=("--steward-key=")
    two_word_flags+=("--steward-key")
    local_nonpersistent_flags+=("--steward-key=")
    flags+=("--steward-name=")
    two_word_flags+=("--steward-name")
    local_nonpersistent_flags+=("--steward-name=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_flag+=("--steward-key=")
    must_have_one_flag+=("--steward-name=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_agency()
{
    last_command="findy-cli_agency"

    command_aliases=()

    commands=()
    commands+=("ping")
    commands+=("start")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_completion()
{
    last_command="findy-cli_completion"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--help")
    flags+=("-h")
    local_nonpersistent_flags+=("--help")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_pool_create()
{
    last_command="findy-cli_pool_create"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--pool-genesis=")
    two_word_flags+=("--pool-genesis")
    local_nonpersistent_flags+=("--pool-genesis=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--poolname=")
    two_word_flags+=("--poolname")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_flag+=("--pool-genesis=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_pool_ping()
{
    last_command="findy-cli_pool_ping"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--poolname=")
    two_word_flags+=("--poolname")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_pool()
{
    last_command="findy-cli_pool"

    command_aliases=()

    commands=()
    commands+=("create")
    commands+=("ping")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--poolname=")
    two_word_flags+=("--poolname")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_flag+=("--poolname=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_service_connect()
{
    last_command="findy-cli_service_connect"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--pwendp=")
    two_word_flags+=("--pwendp")
    local_nonpersistent_flags+=("--pwendp=")
    flags+=("--pwkey=")
    two_word_flags+=("--pwkey")
    local_nonpersistent_flags+=("--pwkey=")
    flags+=("--pwname=")
    two_word_flags+=("--pwname")
    local_nonpersistent_flags+=("--pwname=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_service_createkey()
{
    last_command="findy-cli_service_createkey"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--seed=")
    two_word_flags+=("--seed")
    local_nonpersistent_flags+=("--seed=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--seed=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_service_creddef_create()
{
    last_command="findy-cli_service_creddef_create"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--tag=")
    two_word_flags+=("--tag")
    local_nonpersistent_flags+=("--tag=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--schema-id=")
    two_word_flags+=("--schema-id")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--tag=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_service_creddef_read()
{
    last_command="findy-cli_service_creddef_read"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--creddef-id=")
    two_word_flags+=("--creddef-id")
    local_nonpersistent_flags+=("--creddef-id=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--schema-id=")
    two_word_flags+=("--schema-id")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--creddef-id=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_service_creddef()
{
    last_command="findy-cli_service_creddef"

    command_aliases=()

    commands=()
    commands+=("create")
    commands+=("read")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--schema-id=")
    two_word_flags+=("--schema-id")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--schema-id=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_service_export()
{
    last_command="findy-cli_service_export"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--export-file=")
    two_word_flags+=("--export-file")
    local_nonpersistent_flags+=("--export-file=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--export-file=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_service_invitation()
{
    last_command="findy-cli_service_invitation"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--label=")
    two_word_flags+=("--label")
    local_nonpersistent_flags+=("--label=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--label=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_service_onboard()
{
    last_command="findy-cli_service_onboard"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--email=")
    two_word_flags+=("--email")
    local_nonpersistent_flags+=("--email=")
    flags+=("--export-file=")
    two_word_flags+=("--export-file")
    local_nonpersistent_flags+=("--export-file=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--email=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_service_ping()
{
    last_command="findy-cli_service_ping"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--sa")
    flags+=("-s")
    local_nonpersistent_flags+=("--sa")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_service_schema_create()
{
    last_command="findy-cli_service_schema_create"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--schema-attrs=")
    two_word_flags+=("--schema-attrs")
    local_nonpersistent_flags+=("--schema-attrs=")
    flags+=("--schema-name=")
    two_word_flags+=("--schema-name")
    local_nonpersistent_flags+=("--schema-name=")
    flags+=("--schema-v=")
    two_word_flags+=("--schema-v")
    local_nonpersistent_flags+=("--schema-v=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--schema-name=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_service_schema_read()
{
    last_command="findy-cli_service_schema_read"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--schema-id=")
    two_word_flags+=("--schema-id")
    local_nonpersistent_flags+=("--schema-id=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--schema-id=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_service_schema()
{
    last_command="findy-cli_service_schema"

    command_aliases=()

    commands=()
    commands+=("create")
    commands+=("read")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_service_send()
{
    last_command="findy-cli_service_send"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--con-id=")
    two_word_flags+=("--con-id")
    local_nonpersistent_flags+=("--con-id=")
    flags+=("--from=")
    two_word_flags+=("--from")
    local_nonpersistent_flags+=("--from=")
    flags+=("--msg=")
    two_word_flags+=("--msg")
    local_nonpersistent_flags+=("--msg=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--con-id=")
    must_have_one_flag+=("--msg=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_service_steward()
{
    last_command="findy-cli_service_steward"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--poolname=")
    two_word_flags+=("--poolname")
    local_nonpersistent_flags+=("--poolname=")
    flags+=("--steward-seed=")
    two_word_flags+=("--steward-seed")
    local_nonpersistent_flags+=("--steward-seed=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_service_trustping()
{
    last_command="findy-cli_service_trustping"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--con-id=")
    two_word_flags+=("--con-id")
    local_nonpersistent_flags+=("--con-id=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--con-id=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_service()
{
    last_command="findy-cli_service"

    command_aliases=()

    commands=()
    commands+=("connect")
    commands+=("createkey")
    commands+=("creddef")
    commands+=("export")
    commands+=("invitation")
    commands+=("onboard")
    commands+=("ping")
    commands+=("schema")
    commands+=("send")
    commands+=("steward")
    commands+=("trustping")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_flag+=("--walletkey=")
    must_have_one_flag+=("--walletname=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_user_connect()
{
    last_command="findy-cli_user_connect"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--pwendp=")
    two_word_flags+=("--pwendp")
    local_nonpersistent_flags+=("--pwendp=")
    flags+=("--pwkey=")
    two_word_flags+=("--pwkey")
    local_nonpersistent_flags+=("--pwkey=")
    flags+=("--pwname=")
    two_word_flags+=("--pwname")
    local_nonpersistent_flags+=("--pwname=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_user_createkey()
{
    last_command="findy-cli_user_createkey"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--seed=")
    two_word_flags+=("--seed")
    local_nonpersistent_flags+=("--seed=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--seed=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_user_creddef_read()
{
    last_command="findy-cli_user_creddef_read"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--creddef-id=")
    two_word_flags+=("--creddef-id")
    local_nonpersistent_flags+=("--creddef-id=")
    flags+=("--schema-id=")
    two_word_flags+=("--schema-id")
    local_nonpersistent_flags+=("--schema-id=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--creddef-id=")
    must_have_one_flag+=("--schema-id=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_user_creddef()
{
    last_command="findy-cli_user_creddef"

    command_aliases=()

    commands=()
    commands+=("read")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_user_export()
{
    last_command="findy-cli_user_export"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--export-file=")
    two_word_flags+=("--export-file")
    local_nonpersistent_flags+=("--export-file=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--export-file=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_user_invitation()
{
    last_command="findy-cli_user_invitation"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--label=")
    two_word_flags+=("--label")
    local_nonpersistent_flags+=("--label=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--label=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_user_onboard()
{
    last_command="findy-cli_user_onboard"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--email=")
    two_word_flags+=("--email")
    local_nonpersistent_flags+=("--email=")
    flags+=("--export-file=")
    two_word_flags+=("--export-file")
    local_nonpersistent_flags+=("--export-file=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--email=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_user_ping()
{
    last_command="findy-cli_user_ping"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--sa")
    flags+=("-s")
    local_nonpersistent_flags+=("--sa")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_user_schema_read()
{
    last_command="findy-cli_user_schema_read"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--schema-id=")
    two_word_flags+=("--schema-id")
    local_nonpersistent_flags+=("--schema-id=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--schema-id=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_user_schema()
{
    last_command="findy-cli_user_schema"

    command_aliases=()

    commands=()
    commands+=("read")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_user_send()
{
    last_command="findy-cli_user_send"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--con-id=")
    two_word_flags+=("--con-id")
    local_nonpersistent_flags+=("--con-id=")
    flags+=("--from=")
    two_word_flags+=("--from")
    local_nonpersistent_flags+=("--from=")
    flags+=("--msg=")
    two_word_flags+=("--msg")
    local_nonpersistent_flags+=("--msg=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--con-id=")
    must_have_one_flag+=("--msg=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_user_trustping()
{
    last_command="findy-cli_user_trustping"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--con-id=")
    two_word_flags+=("--con-id")
    local_nonpersistent_flags+=("--con-id=")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")

    must_have_one_flag=()
    must_have_one_flag+=("--con-id=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_user()
{
    last_command="findy-cli_user"

    command_aliases=()

    commands=()
    commands+=("connect")
    commands+=("createkey")
    commands+=("creddef")
    commands+=("export")
    commands+=("invitation")
    commands+=("onboard")
    commands+=("ping")
    commands+=("schema")
    commands+=("send")
    commands+=("trustping")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--url=")
    two_word_flags+=("--url")
    flags+=("--walletkey=")
    two_word_flags+=("--walletkey")
    flags+=("--walletname=")
    two_word_flags+=("--walletname")
    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_flag+=("--walletkey=")
    must_have_one_flag+=("--walletname=")
    must_have_one_noun=()
    noun_aliases=()
}

_findy-cli_root_command()
{
    last_command="findy-cli"

    command_aliases=()

    commands=()
    commands+=("agency")
    commands+=("completion")
    commands+=("pool")
    commands+=("service")
    commands+=("user")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--apiurl=")
    two_word_flags+=("--apiurl")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--data=")
    two_word_flags+=("--data")
    flags+=("--dry-run")
    flags+=("-n")
    flags+=("--logging=")
    two_word_flags+=("--logging")
    flags+=("--salt=")
    two_word_flags+=("--salt")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

__start_findy-cli()
{
    local cur prev words cword
    declare -A flaghash 2>/dev/null || :
    declare -A aliashash 2>/dev/null || :
    if declare -F _init_completion >/dev/null 2>&1; then
        _init_completion -s || return
    else
        __findy-cli_init_completion -n "=" || return
    fi

    local c=0
    local flags=()
    local two_word_flags=()
    local local_nonpersistent_flags=()
    local flags_with_completion=()
    local flags_completion=()
    local commands=("findy-cli")
    local must_have_one_flag=()
    local must_have_one_noun=()
    local last_command
    local nouns=()

    __findy-cli_handle_word
}

if [[ $(type -t compopt) = "builtin" ]]; then
    complete -o default -F __start_findy-cli findy-cli
else
    complete -o default -o nospace -F __start_findy-cli findy-cli
fi

# ex: ts=4 sw=4 et filetype=sh
