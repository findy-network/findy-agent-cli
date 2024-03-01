-- TODO: REMOVE me
local function getRegValue(reg, key)
    _, _ = reg, key
    return ""
end

local function setRegValue(reg, key, val)
    _, _, _ = reg, val, key
end

local function cmd_argument(line)
    local i=string.find(line, ' ')
    if i == nil then
        return line, ""
    end
    local cmd = string.sub(line, 1, i-1)
    local name = string.sub(line, i+1, string.len(line))
    return cmd, name
end

local cmd = getRegValue("MEM", "INPUT")
--local line = getRegValue("MEM", "INPUT")
--local cmd, name = cmd_argument(line)

if cmd == 'issuer' then
    setRegValue('MEM', 'TARGET', 'WAIT_AS_ISSUER')
    setRegValue('MEM', 'OUTPUT', 'OK')
elseif cmd == 'receiver' then
    setRegValue('MEM', 'TARGET', 'WAIT_AS_RECEIVER')
    setRegValue('MEM', 'OUTPUT', 'OK')
else
    setRegValue('MEM', 'TARGET', 'WAIT_ROLE')
    setRegValue('MEM', 'OUTPUT', 'ERR')
end

