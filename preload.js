// Copyright (C) 2024-2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Mestre da Info
// Site: https://www.mestredainfo.com.br

const { contextBridge, ipcRenderer } = require('electron')

ipcRenderer.setMaxListeners(20);

contextBridge.exposeInMainWorld('micheckhash', {
    version: (type) => ipcRenderer.invoke('appVersao', type),
    newWindow: (url, width, height, resizable, frame, menu, hide) => ipcRenderer.invoke('appNewWindow', url, width, height, resizable, frame, menu, hide),
    openURL: (url) => ipcRenderer.invoke('appExterno', url),
    translate: (text, ...values) => ipcRenderer.invoke('appTraduzir', text, ...values),
    openFile: () => ipcRenderer.invoke('appAbrirArquivo'),
    readFile: (filename) => ipcRenderer.invoke('appReadFile', filename),
    getHash: (tipo, arquivo) => ipcRenderer.invoke('appGetHash', tipo, arquivo),
    saveHash: (filename, tipo, hash) => ipcRenderer.invoke('appSaveHash', filename, tipo, hash)
});