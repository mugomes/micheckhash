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

async function checkHash(e) {
    let sTipoHash = document.getElementById('txtTipoHash').value;

    if (document.getElementById('txtArquivo').value == '') {
        document.getElementById('resultado').innerHTML = '<div class="alert alert-danger">Selecione um arquivo para verificar o hash.</div>';
        return false;
    }

    micheckhash.translate('Checking hash...').then((sValue) => {
        document.getElementById('resultado').innerHTML = `<div class="alert alert-info">${sValue}</div>`;
    });

    micheckhash.getHash(sTipoHash, document.getElementById('txtArquivo').value).then((sHashFile) => {
        let sHash = document.getElementById('txtHash').value;

        document.getElementById('hash').innerHTML = sHashFile;

        if (sHashFile == sHash) {
            micheckhash.translate('Success! The file hash is the same as the entered hash.').then((sValue) => {
                document.getElementById('resultado').innerHTML = `<div class="alert alert-success">${sValue}</div>`;
            });
        } else {
            micheckhash.translate('Danger! The file hash is different from the reported hash.').then((sValue) => {
                document.getElementById('resultado').innerHTML = `<div class="alert alert-danger">${sValue}</div>`;
            });
        }
    });

    e.preventDefault();
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

micheckhash.translate(document.getElementById('lblHash').innerHTML).then((sValue) => {
    document.getElementById('lblHash').innerHTML = sValue;
});

micheckhash.translate(document.getElementById('btnCheck').innerHTML).then((sValue) => {
    document.getElementById('btnCheck').innerHTML = sValue;
});