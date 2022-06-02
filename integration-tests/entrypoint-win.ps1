$file = $Env:windir+'\System32\drivers\etc\hosts'; $Env:SAMPLE_IP+' sample' | Add-Content -PassThru $file;

# this will validate that we have a proper set in the hosts file and access the sample ip
$connection = New-Object System.Net.Sockets.TcpClient("sample", 8000)
if ($connection.Connected) {
    go test -v -mod=vendor . --client-id $Env:CLIENT_ID --client-secret $Env:CLIENT_SECRET --sample-url $Env:SAMPLE_URL --redirect-url $Env:REDIRECT_URL --sample-name $Env.SAMPLE_NAME --skip-modify-tests $Env:SKIP_MODIFY_TESTS
    exit $LASTEXITCODE
} else {
    Write-Host "Tests cannot be run as the sample's ip $Env:SAMPLE_IP is not accessible"
    exit 1
}
