param (
  [string]$Version
)

$currentPath = Get-Location
$currentDir = "$currentPath\$Version\output"
$signtoolPath = 'C:\Program Files (x86)\Windows Kits\10\bin\10.0.18362.0\x64\signtool.exe'
$timestampUrl = "http://timestamp.digicert.com"
$certificateThumbprint = $env:CERTIFICATE_THUMBPRINT

Write-Host "Version: $Version"
Write-Host "Current Path: $currentPath"
Write-Host "Current Dir: $currentDir"

& $signtoolPath sign /sha1 "$certificateThumbprint" /tr $timestampUrl /td sha256 /fd sha256 $currentDir\olares-cli_windows_amd64_v1\olares-cli.exe
& $signtoolPath sign /sha1 "$certificateThumbprint" /tr $timestampUrl /td sha256 /fd sha256 $currentDir\olares-cli_windows_arm64\olares-cli.exe

$fileName_x86 = "olares-cli-v{0}_windows_amd64.zip" -f $Version
$fileName_arm64 = "olares-cli-v{0}_windows_arm64.zip" -f $Version

Compress-Archive -Path $currentDir\olares-cli_windows_amd64_v1\olares-cli.exe -DestinationPath "$currentDir\$fileName_x86"
Compress-Archive -Path $currentDir\olares-cli_windows_arm64\olares-cli.exe -DestinationPath "$currentDir\$fileName_arm64"