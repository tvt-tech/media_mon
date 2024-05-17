# Archer USB File Filter

### tool to find connected Archer devices and archive all data not matched to extensions patterns list 

### Build
```powershell
git clone https://github.com/tvt-tech/usb-file-filter
cd usb-file-filter
$env:GOOS = "windows"
$env:GOARCH = "386"
go build -ldflags="-s -w" -trimpath
```

### Usage

Find Archer devices in usb storages and archive founded unexpected files
```powerhsell
usb-file-filter.exe 
```

Find Archer compatiple files signature in specified path
 ```powerhsell
usb-file-filter.exe ./<dst>
```

List USB drives
 ```powerhsell
usb-file-filter.exe -l
```

Run in debug mode
 ```powerhsell
usb-file-filter.exe -d ./<dst>
```