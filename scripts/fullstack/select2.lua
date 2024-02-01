local function cmd_argument(line)
    local i=string.find(line, ' ')
    if i == nil then
        return line, ""
    end
    local cmd = string.sub(line, 1, i-1)
    local name = string.sub(line, i+1, string.len(line))
    return cmd, name
end

local line = getRegValue("MEM", "INPUT")
local cmd, name = cmd_argument(line)

if cmd == 'name' then
    setRegValue('MEM', 'TARGET', 'RUN')
    setRegValue('MEM', 'NAME', name)
    setRegValue('MEM', 'OUTPUT', 'OK')
elseif cmd == 'help' then
    setRegValue('MEM', 'TARGET', 'IDLE')
    setRegValue('MEM', 'OUTPUT', 'OK')
elseif cmd == 'quit' then
    setRegValue('MEM', 'TARGET', 'QUIT')
    setRegValue('MEM', 'OUTPUT', 'OK')
end
