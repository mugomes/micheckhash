// Copyright (C) 2024-2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Mestre da Info
// Site: https://www.mestredainfo.com.br

const { contextBridge, ipcRenderer } = require('electron')

ipcRenderer.setMaxListeners(20);

contextBridge.exposeInMainWorld('micheckhash', {
    version: (type) => ipcRenderer.invoke('appVersao', type),
    alert: (title, msg, type, button) => ipcRenderer.invoke('appMessage', title, msg, type, button),
    confirm: (title, msg, type, ...buttons) => ipcRenderer.invoke('appConfirm', title, msg, type, ...buttons),
    newWindow: (url, width, height, resizable, frame, menu, hide) => ipcRenderer.invoke('appNewWindow', url, width, height, resizable, frame, menu, hide),
    openURL: (url) => ipcRenderer.invoke('appExterno', url),
    translate: (text, ...values) => ipcRenderer.invoke('appTraduzir', text, ...values),
    selectDirectory: () => ipcRenderer.invoke('appSelecionarDiretorio'),
    openFile: () => ipcRenderer.invoke('appAbrirArquivo'),
    saveFile: () => ipcRenderer.invoke('appSalvarArquivo'),
    readFile: (filename) => ipcRenderer.invoke('appReadFile', filename),
    notification: (title, text) => ipcRenderer.invoke('appNotification', title, text),
    tray: (title, tooltip, icon, menu) => ipcRenderer.invoke('appTray', title, tooltip, icon, menu),
    exportPDF: (filename, options) => ipcRenderer.invoke('appExportPDF', filename, options),
    devTools: () => ipcRenderer.invoke('appDevTools'),
    close: () => ipcRenderer.invoke('appSair'),
    getHash: (tipo, arquivo) => ipcRenderer.invoke('appGetHash', tipo, arquivo),
    saveHash: (filename, tipo, hash) => ipcRenderer.invoke('appSaveHash', filename, tipo, hash)
});