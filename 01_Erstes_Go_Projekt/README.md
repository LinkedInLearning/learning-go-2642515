
# Erstes Go Projekt

## Ihr erstes Go Projekt mit Modulen und Packages

Go Modul erstellen mit `go mod init linkedin.com/gopro`.


## Mit den Go Tools arbeiten

### 1. Go Build

```bash
$ go help build        # Hilfe anzeigen
$ go build .           # Code kompilieren
$ go build -o app .    # Code kompilieren mit Binary Output app
```

### 2. Go Format

```bash
$ go fmt main.go   # Eine Datei formatieren
$ go fmt .         # Alle Dateien des aktuelle Verzeichnisses formatieren

### 3. Go Doc

```bash
$ go doc strings            # Anzeige der Dokumentation des strings Packages
$ go doc -all strings       # Anzeige der ausführlichen Dokumentation des strings Packages
$ go doc strings.Replace    # Anzeige der Dokumentation der strings.Replace Funktion
```
