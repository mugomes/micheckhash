# Copyright (C) 2024-2025 Murilo Gomes Julio
# SPDX-License-Identifier: GPL-2.0-only

# Site: https://github.com/mugomes

import tkinter as tk
import ttkbootstrap as ttk
import os, hashlib

import translate as cTrans

from tkinter import Frame, StringVar, Menu, messagebox
from tkinter.filedialog import askopenfilename, asksaveasfilename
from ttkbootstrap.constants import *

def showWindow():
    _ = cTrans.getLang()
    
    TAMANHO_BLOCO = 65536
    FILE_EXTENSION = '.md5'

    def saveHash():
        nonlocal txtFileVar, txtCodigoVar, FILE_EXTENSION

        sFileName = os.path.basename(txtFileVar.get())
        sCodigo = txtCodigoVar.get()

        file_path = asksaveasfilename(
            defaultextension=f'{FILE_EXTENSION}',
            filetypes=[(_('Text File'), f'*{FILE_EXTENSION}')],
            title=_('Save As')
        )

        if file_path:
            with open(file_path, "w", encoding="utf-8") as f:
                f.write(f'{sCodigo} {sFileName}')

    def openDialog():
        nonlocal txtFileVar
        sValue = askopenfilename(
            title=_('Open File')
        )
        
        if sValue:
            txtFileVar.set(sValue)

    def copyCodigo():
        nonlocal txtCodigoVar
        try:
            frmGerar.clipboard_append(string=txtCodigoVar.get())
        except:
            print('')

    def gerarHash():
        nonlocal txtFileVar, cboTipoHashVar, txtCodigoVar, TAMANHO_BLOCO, FILE_EXTENSION
        try:
            arquivoSelecionado = txtFileVar.get()
            
            if not os.path.exists(arquivoSelecionado):
                return None
        
            tipohash = hashlib.md5
            if cboTipoHashVar.get() == 'SHA1':
                tipohash = hashlib.sha1
                FILE_EXTENSION = '.sha1'
            elif cboTipoHash.get() == 'SHA256':
                tipohash = hashlib.sha256
                FILE_EXTENSION = '.sha256'
            elif cboTipoHash.get() == 'SHA512':
                tipohash = hashlib.sha512
                FILE_EXTENSION = '.sha512'
            
            hasher = tipohash()
            
            with open(arquivoSelecionado, 'rb') as f:
                buf = f.read(TAMANHO_BLOCO)
                while len(buf) > 0:
                    hasher.update(buf)
                    buf = f.read(TAMANHO_BLOCO)
            
            # Retorna o hash hexadecimal
            resultado = hasher.hexdigest()

            if resultado is None:
                messagebox.showinfo(title='MiCheckHash', message=_('Could not generate file hash.'), icon="error")
                return
            
            txtCodigoVar.set(resultado)
        except Exception as e:
            messagebox.showinfo(title='MiCheckHash', message=_('Error reading file: %s') % {e}, icon="error")
            return None

    # Janela Principal
    frmGerar = ttk.Toplevel()
    frmGerar.title(_('Generate - %s') % 'MiCheckHash')
    frmGerar.geometry('400x310')
    frmGerar.position_center()
    frmGerar.resizable(False, False)

    icon_image = tk.PhotoImage(file=r'micheckhash.png')
    frmGerar.iconphoto(True, icon_image) 

    # Estilos
    style = ttk.Style()
    style.configure('btnCheck.TButton', font=('', 12, 'bold'))

    # Menus
    barmenuMain = Menu(frmGerar)
    frmGerar.config(menu=barmenuMain)

    mnuFiles = Menu(barmenuMain, tearoff=0)
    barmenuMain.add_cascade(label=_('File'), menu=mnuFiles)
    mnuFiles.add_command(label=_('Save Hash'), command=saveHash)

    cboTipoHashVar = StringVar(value='MD5')
    lblTipoHash = ttk.Label(frmGerar, text=_('Select hash type'), font=("", 12, "bold"))
    lblTipoHash.pack(anchor="nw", padx=5, pady=(7, 3))
    cboTipoHash = ttk.Combobox(frmGerar, textvariable=cboTipoHashVar, values=['MD5', 'SHA1', 'SHA256', 'SHA512'], bootstyle="info", state='readonly')
    cboTipoHash.pack(anchor="nw", padx=5, fill='x')

    txtFileVar = StringVar(value='')
    lblSelectFile = ttk.Label(frmGerar, text=_('Select file'), font=('', 12, 'bold'))
    lblSelectFile.pack(anchor='nw', padx=5, pady=(17, 3))
    frameFile = Frame(frmGerar)
    frameFile.pack(anchor='nw', fill='x')
    txtFile = ttk.Entry(frameFile, bootstyle="info", textvariable=txtFileVar, font=('', 12), state=READONLY)
    txtFile.pack(fill='x', padx=5, side=LEFT, expand=True)
    btnEllipsisFile = ttk.Button(frameFile, text='...', width=3, bootstyle='secondary', command=openDialog)
    btnEllipsisFile.pack(anchor='nw', padx=5, side=LEFT)

    txtCodigoVar = StringVar()
    lblCodigo = ttk.Label(frmGerar, text=_('Generated Code'), font=('', 12, 'bold'))
    lblCodigo.pack(anchor='nw', padx=5, pady=(17, 3))
    frameCodigo = Frame(frmGerar)
    frameCodigo.pack(anchor='nw', fill='x')
    txtCodigo = ttk.Entry(frameCodigo, bootstyle="info", textvariable=txtCodigoVar, font=('', 12))
    txtCodigo.pack(fill='x', padx=5, side=LEFT, expand=True)
    btnCopiarCodigo = ttk.Button(frameCodigo, text=_('Copy'), command=copyCodigo, bootstyle='secondary')
    btnCopiarCodigo.pack(anchor='nw', padx=5, side=LEFT)

    btnCheck = ttk.Button(frmGerar, text=_('Generate Hash'), style='btnCheck.TButton', command=gerarHash)
    btnCheck.pack(pady=(17, 5))

    # Carrega Janela
    return frmGerar