#!/bin/bash

if [[ $_ == $0 || $1 == "" ]]; then
	printf "Usage:\tsource %s <name>\n" "$0"

	exit 1
fi

recover_file="./recover-names.sh"

if [[ -f "$recover_file" ]]; then
	rm "$recover_file"
fi
echo "#!/bin/bash" > "$recover_file"
cat > "$recover_file" <<EOF
#!/bin/bash 

if [[ "\$_" == "\$0" ]]; then
	printf "Usage:\tsource %s\n" "$0"
	exit 1
fi

# your agent names:
EOF

for a in "$@"; do
	uuid=`uuidgen`
	cmd="export $a=$a-$uuid"
	echo "$cmd" >> "$recover_file"
	eval "$cmd"
	printf "your variable \$%s is ready for use\n" "$a"
done
