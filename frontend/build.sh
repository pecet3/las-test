#!/bin/bash

# Krok 1: Uruchomienie npm run build
echo "Uruchamiam npm run build..."
npm run build

# Sprawdzenie, czy build się powiódł
if [ $? -eq 0 ]; then
  echo "Build zakończony pomyślnie."
else
  echo "Błąd podczas budowania aplikacji."
  exit 1
fi

# Krok 2: Usuwanie istniejącego folderu view w ../backend/cmd, jeśli istnieje
if [ -d "../backend/cmd/view" ]; then
  echo "Folder view już istnieje, usuwam go..."
  rm -rf ../backend/cmd/view
  # Sprawdzamy, czy usunięcie się powiodło
  if [ $? -eq 0 ]; then
    echo "Folder view został usunięty."
  else
    echo "Błąd podczas usuwania folderu view."
    exit 1
  fi
fi

# Krok 3: Przenoszenie folderu dist do ../backend/cmd/view
echo "Przenoszę folder dist do ../backend/cmd/view..."
mv dist ../backend/cmd/view

# Sprawdzenie, czy przenoszenie się powiodło
if [ $? -eq 0 ]; then
  echo "Folder dist został przeniesiony pomyślnie jako view."
else
  echo "Błąd podczas przenoszenia folderu dist."
  exit 1
fi

echo "Proces zakończony."
