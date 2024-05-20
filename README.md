# Archer USB File Filter

### tool to find connected Archer devices and archive all data not matched to extensions patterns list 

### Build
WINDOWS x32/amd64
```powershell
git clone https://github.com/tvt-tech/archer_media_mon
cd usb-file-filter
$env:GOOS = "windows"
$env:GOARCH = "386"
go build -ldflags="-s -w" -trimpath -o media_mon_x32.exe
```

LINUX
```bash
git clone https://github.com/tvt-tech/usb-file-filter
cd usb-file-filter
GOOS=linux go build -ldflags="-s -w" -trimpath -o media_mon
```

MIPSLE
```bash
git clone https://github.com/tvt-tech/usb-file-filter
cd usb-file-filter
GOOS=linux GOARCH=mipsle CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o media_mon
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
```bash
./media_mon ./<dst>
```

List USB drives
```bash
./media_mon -l
```

Eject drive by path 
```bash
./media_mon -e <drive path / drive letter>
```

Eject all matched Archer devices
```bash
./media_mon -e -A
```

Run in monitor mode
```bash
./media_mon -s
```

Run in monitor mode with no tray icon
```bash
./media_mon -s -q
```

Run in debug mode
```bash
./media_mon -d ./<dst>
```