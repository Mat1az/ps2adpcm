> [!IMPORTANT]
> Alpha release: Intended for use with GTA: San Andreas SFX files; not yet tested on other games.

<h1 align="center">
  <br>
  PS2ADPCM
  <br>
</h1>

[![License: GPL-3.0](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0.html)

<h4 align="center">A Minimal Audio Converter for PlayStation 2 (PS2) Audio Files</h4>


<p align="center">
  <a href="#how-it-works">How It Works</a> •
  <a href="#key-features">Key Features</a> •
  <a href="#how-to-use">How To Use</a> •
  <a href="#download">Download</a> •
  <a href="#how-to-do-it-manually">How To Do It Manually</a>
</p>

## How It Works
![PS2ADPCM](https://github.com/Mat1az/ps2adpcm/assets/39683242/66a4ff7e-36ef-41d8-807c-b9db8618337b)


## Key Features

* Simple functionality
  - Automatically generates a valid audio file ready for import into a PlayStation 2 (PS2) game.

## How To Use
```console
./ps2adpcm
```
![Screenshot_2023-11-27_12-26-48](https://github.com/Mat1az/ps2adpcm/assets/39683242/69dfae49-ccda-44fb-9a20-0beb457b13fb)

 OR
```console
./ps2adpcm -c [custom audio] -i [original audio] -o [output audio]
```

## Download

You can [download](https://github.com/Mat1az/ps2adpcm/releases/latest) the latest version of PS2ADPCM.

## How To Do It Manually
1. Open the original WAV file using MFAudio or Audacity.
   - Copy the `Frequency Hz` and paste it into Notepad.
2. Open the original WAV file with a Hex Editor.
   - Copy the `first 44 HEX values` and paste them into Notepad.
   - Copy the `last 16 HEX values` and paste them into Notepad.
3. Open your custom WAV file with Audacity.
   - Go to Edit > Preferences > Audio Settings > Quality.
   - Set the Project Sample Rate to `paste the Frequency Hz from your original WAV`.
   - Make sure the duration of your custom WAV file is the same as or shorter than the original WAV.
   - Save and Replace.
4. Open the custom WAV file with MFAudio.
   - Convert it into `RAW - Raw Sound Data - Compressed ADPCM`.
5. Open the custom WAV file with a Hex Editor.
   - Insert the original `first 44 HEX values` at the beginning of the custom WAV.
   - Remove the `last 32 HEX values` from the end of the custom WAV.
   - Insert the original `last 16 HEX values` at the end of the custom WAV.
