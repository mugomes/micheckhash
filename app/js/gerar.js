// Copyright (C) 2024-2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Mestre da Info
// Site: https://www.mestredainfo.com.br

document.getElementById('txtArquivo').addEventListener('click', async () => {
    let sFilename = await micheckhash.openFile();

    if (typeof sFilename !== 'undefined') {
        document.getElementById('txtArquivo').value = sFilename;
    }
});

async function generateHash(e) {
    let sTipoHash = document.getElementById('txtTipoHash').value;

    if (document.getElementById('txtArquivo').value == '') {
        document.getElementById('resultado').innerHTML = '<div class="alert alert-danger">Selecione um arquivo para gerar o hash.</div>';
        return false;
    }

    micheckhash.translate('Generating hash...').then((sValue) => {
        document.getElementById('resultado').innerHTML = `<div class="alert alert-info">${sValue}</div>`;
    });

    micheckhash.getHash(sTipoHash, document.getElementById('txtArquivo').value).then((sHashFile) => {
        SalvarArquivo(document.getElementById('txtArquivo').value, document.getElementById('txtTipoHash').value, sHashFile);
    });

    e.preventDefault();
}

// Salvar Arquivo
async function SalvarArquivo(filename, tipo, hash) {
    await micheckhash.saveHash(filename, tipo, hash).then((value) => {
        if (value) {
            micheckhash.translate('Hash saved successfully!').then((sValue) => {
                document.getElementById('resultado').innerHTML = `<div class="alert alert-success">${sValue}</div>`;
            });
        } else {
            micheckhash.translate('An error occurred, unable to save the hash!').then((sValue) => {
                document.getElementById('resultado').innerHTML = `<div class="alert alert-danger">${sValue}</div>`;
            });

        }
    });
}

micheckhash.translate(document.getElementById('lblTipoHash').innerHTML).then((sValue) => {
    document.getElementById('lblTipoHash').innerHTML = sValue;
});

micheckhash.translate(document.getElementById('lblArquivo').innerHTML).then((sValue) => {
    document.getElementById('lblArquivo').innerHTML = sValue;
});

micheckhash.translate(document.getElementById('txtArquivo').getAttribute('placeholder')).then((sValue) => {
    document.getElementById('txtArquivo').setAttribute('placeholder', sValue);
});

micheckhash.translate(document.getElementById('btnCheck').innerHTML).then((sValue) => {
    document.getElementById('btnCheck').innerHTML = sValue;
});