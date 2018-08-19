# ApiGo (https://github.com/eGenGuru/ApiGo)
# -----------------------------------------
# script sample for Powershell

function HelloWorld {
    param (
    )
    @{ "message" = "Hello World!" } | ConvertTo-Json
}

Export-ModuleMember -Function HelloWorld
