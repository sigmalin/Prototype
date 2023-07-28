set location=%cd%
set input=src
set output=grpc
set cmd=%location%\protoc --proto_path=./%input% --csharp_out=./%output% --grpc_out=./%output% --plugin=protoc-gen-grpc="%cd%\grpc_csharp_plugin.exe"

if not exist %input% (mkdir %input%)
if not exist %output% (mkdir %output%)

for /F %%i IN ('dir %input%\*.proto /b') DO %cmd% %%i