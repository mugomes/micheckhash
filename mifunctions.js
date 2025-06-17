// Copyright (C) 2024-2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Mestre da Info
// Site: https://www.mestredainfo.com.br

const { ipcMain, dialog, BrowserWindow } = require('electron')

module.exports = {
    mifunctions: function (win, milang, miNewWindow, miPath) {
        // Função para abrir arquivo
        ipcMain.handle('appAbrirArquivo', async () => {
            const { canceled, filePaths } = await dialog.showOpenDialog({ properties: ['openFile'] });
            if (!canceled) {
                return filePaths[0];
            }
        });

        // Abrir aplicativo externo
        ipcMain.handle('appExterno', async (event, url) => {
            require('electron').shell.openExternal(url);
        });

        // Obter versão do aplicativo e recursos
        ipcMain.handle('appVersao', async (event, tipo) => {
            if (tipo == 'micheckhash') {
                return require('electron').app.getVersion();
            } else if (tipo == 'electron') {
                return process.versions.electron;
            } else if (tipo == 'node') {
                return process.versions.node;
            } else if (tipo == 'chromium') {
                return process.versions.chrome;
            } else {
                return '';
            }
        });

        // Abre uma nova janela personalizada
        ipcMain.handle('appNewWindow', async (event, url, width, height, resizable, frame, menu, hide) => {
            miNewWindow(url, width, height, resizable, frame, menu, hide);
        });

        // Traduzir
        ipcMain.handle('appTraduzir', async (event, text, ...values) => {
            return milang.traduzir(text, ...values);
        });

        // Função para ler arquivo
        ipcMain.handle('appReadFile', async (event, filename, externo) => {
            const fs = require('fs');
            const path = require('path');
            try {
                if (externo) {
                    return fs.readFileSync(filename, "utf8");
                } else {
                    return fs.readFileSync(path.join(miPath, filename), "utf8");
                }
            } catch (err) {
                return false;
            }
        });

        ipcMain.handle('appSaveHash', async (event, filename, tipo, hash) => {
            const fs = require('fs');
            const path = require('path');

            try {
                if (fs.existsSync(filename)) {
                    data = hash + ' ' + path.basename(filename);
                    fs.writeFileSync(`${filename}.${tipo}`, data);
                } else {
                    return false;
                }

                return true;
            } catch (err) {
                return false;
            }
        });

        // Terminal
        ipcMain.handle('appGetHash', async (event, tipo, arquivo) => {
            const sCrypto = require('crypto');
            const fs = require('fs');

            try {
                if (fs.existsSync(arquivo)) {
                    let sFilename = fs.createReadStream(arquivo);
                    let sCript = sCrypto.createHash(tipo);

                    sFilename.on('data', (data) => {
                        sCript.update(data);
                    });

                    return new Promise((resolve) => {
                        sFilename.on('end', () => {
                            resolve(sCript.digest('hex'));
                        });
                    });
                } else {
                    return false;
                }
            } catch (err) {
                return err;
            }
        });
    }
}