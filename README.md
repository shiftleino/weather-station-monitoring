# Weather Station Data Pipeline
This repository contains the source code for a Netatmo Weather Station Pipeline. As Netatmo has removed the Client Credentials grant type method, the code in this repository is **obsolete**.

## Workflow
The data is first extracted from the official API of Netatmo (https://api.netatmo.com/). Then the data is transformed to contain only the relevant weather data with right date formats. Finally the data is loaded into Google Sheets using Google's official API (https://developers.google.com/sheets/api). The final data is visualized in Google Data Studio (not shown in this repository).

## Note
This application is not an official production of Netatmo and is for personal, non-commercial use only.
