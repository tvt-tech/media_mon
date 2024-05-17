# Archer USB File Filter

### tool to find connected Archer devices and archive all data not matched to extensions patterns list 

### Build
WINDOWS
```powershell
git clone https://github.com/tvt-tech/usb-file-filter
cd usb-file-filter
$env:GOOS = "windows"
$env:GOARCH = "386"
go build -ldflags="-s -w" -trimpath -o arch_usb_mon_x32.exe
```

LINUX
```bash
git clone https://github.com/tvt-tech/usb-file-filter
cd usb-file-filter
GOOS=linux GOARCH=386 go build -ldflags="-s -w" -trimpath -o arch_usb_mon_386
```

### Usage

Find Archer devices in usb storages 
and archive founded unexpected files
> [!NOTE] Will show popup message on windows to accept or decline action!
```powerhsell
usb-file-filter.exe 
```

Find Archer compatiple files signature in specified path 
and archive founded unexpected files
> [!NOTE] Will show popup message on windows to accept or decline action!
```powerhsell
usb-file-filter.exe ./<dst>
```

List USB drives
```powerhsell
usb-file-filter.exe -l
```

Eject drive by path 
```powerhsell
usb-file-filter.exe -e E:
```

Eject all matched Archer devices
```powerhsell
usb-file-filter.exe -e -A
```

Run as daemon
```powerhsell
usb-file-filter.exe -s
```

Run as daemon with no tray icon
```powerhsell
usb-file-filter.exe -s -q
```

Run in debug mode
```powerhsell
usb-file-filter.exe -d ./<dst>
```