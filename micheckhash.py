# Copyright (C) 2024-2025 Murilo Gomes Julio
# SPDX-License-Identifier: GPL-2.0-only

# Site: https://github.com/mugomes

import tkinter as tk
import ttkbootstrap as ttk
import os, hashlib, threading, webbrowser

# Controls
import translate as cTrans

# Forms
import frmGerar, frmAbout

from tkinter import Frame, StringVar, Menu, messagebox
from tkinter.filedialog import askopenfilename
from ttkbootstrap.constants import *

TAMANHO_BLOCO = 65536

_ = cTrans.getLang()

def gerarHash():
    frmGerar.showWindow()


def checkUpdate():
    webbrowser.open(url="https://github.com/mugomes/micheckhash/releases")


def supportApp():
    webbrowser.open(url="https://github.com/mugomes/micheckhash")


def showAbout():
    frmAbout.showWindow()


def openDialog():
    global txtFileVar
    sValue = askopenfilename(title=_("Open File"))

    if sValue:
        txtFileVar.set(sValue)


def pasteCodigo():
    global txtCodigoVar
    try:
        txtCodigoVar.set(frmMain.clipboard_get())
    except:
        print("")


def checkHash():
    global txtFileVar, cboTipoHashVar, TAMANHO_BLOCO
    try:
        lblInfo.pack()
        btnCheck.config(state="disabled")

        arquivoSelecionado = txtFileVar.get()
        codigoColado = txtCodigoVar.get().strip()

        resultado = ""

        def worker():
            nonlocal resultado
            if not os.path.exists(arquivoSelecionado):
                return None

            tipohash = hashlib.md5
            if cboTipoHashVar.get() == "SHA1":
                tipohash = hashlib.sha1
            elif cboTipoHash.get() == "SHA256":
                tipohash = hashlib.sha256
            elif cboTipoHash.get() == "SHA512":
                tipohash = hashlib.sha512

            hasher = tipohash()

            with open(arquivoSelecionado, "rb") as f:
                buf = f.read(TAMANHO_BLOCO)
                while len(buf) > 0:
                    hasher.update(buf)
                    buf = f.read(TAMANHO_BLOCO)

            # Retorna o hash hexadecimal
            resultado = hasher.hexdigest()

        wThread = threading.Thread(target=worker, daemon=True)
        wThread.start()

        def checkThread():
            if wThread.is_alive():
                frmMain.after(100, checkThread)  # verifica novamente depois de 100ms
            else:
                if resultado is None:
                    messagebox.showinfo(
                        title="MiCheckHash",
                        message=_("Could not generate file hash."),
                        icon="error",
                    )
                    return

                if resultado == codigoColado:
                    messagebox.showinfo(
                        title="MiCheckHash", message=_("Success!"), icon="info"
                    )
                else:
                    messagebox.showinfo(
                        title="MiCheckHash", message=_("Different!"), icon="error"
                    )
                lblInfo.pack_forget()
                btnCheck.config(state="normal")

        checkThread()
    except Exception as e:
        messagebox.showinfo(
            title="MiCheckHash", message=_("Error reading file: %s") % {e}, icon="error"
        )
        lblInfo.pack_forget()
        btnCheck.config(state="normal")
        return None


# Janela Principal
frmMain = ttk.Window(title="MiCheckHash", themename="darkly")
frmMain.geometry("400x350")
frmMain.position_center()
frmMain.resizable(False, False)

icon_image = tk.PhotoImage(file=r'micheckhash.png')
frmMain.iconphoto(True, icon_image) 

# Estilos
style = ttk.Style()
style.configure("btnCheck.TButton", font=("", 12, "bold"))

# Menus
barmenuMain = Menu(frmMain)
frmMain.config(menu=barmenuMain)

mnuTools = Menu(barmenuMain, tearoff=0)
barmenuMain.add_cascade(label=_("Tools"), menu=mnuTools)
mnuTools.add_command(label=_("Generate Hash"), command=gerarHash)

mnuAbout = Menu(barmenuMain, tearoff=0)
barmenuMain.add_cascade(label=_("About"), menu=mnuAbout)
mnuAbout.add_command(label=_("Check Update"), command=checkUpdate)
mnuAbout.add_separator()
mnuAbout.add_command(label=_("Support MiCheckHash"), command=supportApp)
mnuAbout.add_separator()
mnuAbout.add_command(label=_("About MiCheckHash"), command=showAbout)

cboTipoHashVar = StringVar(value="MD5")
lblTipoHash = ttk.Label(frmMain, text=_("Select hash type"), font=("", 12, "bold"))
lblTipoHash.pack(anchor="nw", padx=5, pady=(7, 3))
cboTipoHash = ttk.Combobox(
    frmMain,
    textvariable=cboTipoHashVar,
    values=["MD5", "SHA1", "SHA256", "SHA512"],
    bootstyle="info",
    state="readonly",
)
cboTipoHash.pack(anchor="nw", padx=5, fill="x")

txtFileVar = StringVar(value="")
lblSelectFile = ttk.Label(frmMain, text=_("Select the file"), font=("", 12, "bold"))
lblSelectFile.pack(anchor="nw", padx=5, pady=(17, 3))
frameFile = Frame(frmMain)
frameFile.pack(anchor="nw", fill="x")
txtFile = ttk.Entry(
    frameFile, bootstyle="info", textvariable=txtFileVar, font=("", 12), state=READONLY
)
txtFile.pack(fill="x", padx=5, side=LEFT, expand=True)
btnEllipsisFile = ttk.Button(
    frameFile, text="...", width=3, bootstyle="secondary", command=openDialog
)
btnEllipsisFile.pack(anchor="nw", padx=5, side=LEFT)

txtCodigoVar = StringVar()
lblCodigo = ttk.Label(frmMain, text=_("Type/Paste the Hash"), font=("", 12, "bold"))
lblCodigo.pack(anchor="nw", padx=5, pady=(17, 3))
frameCodigo = Frame(frmMain)
frameCodigo.pack(anchor="nw", fill="x")
txtCodigo = ttk.Entry(
    frameCodigo, bootstyle="info", textvariable=txtCodigoVar, font=("", 12)
)
txtCodigo.pack(fill="x", padx=5, side=LEFT, expand=True)
btnColarCodigo = ttk.Button(
    frameCodigo, text=_("Paste"), command=pasteCodigo, bootstyle="secondary"
)
btnColarCodigo.pack(anchor="nw", padx=5, side=LEFT)

btnCheck = ttk.Button(
    frmMain, text=_("Check Now"), style="btnCheck.TButton", command=checkHash
)
btnCheck.pack(pady=(17, 5))

lblInfo = ttk.Label(
    frmMain, text=_("Verifying Hash... Please wait!"), font=("", 12, "bold")
)
lblInfo.pack(anchor="c", padx=5, pady=(17, 3))
lblInfo.pack_forget()

# Carrega Janela
frmMain.mainloop()
