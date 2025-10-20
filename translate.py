# Copyright (C) 2024-2025 Murilo Gomes Julio
# SPDX-License-Identifier: GPL-2.0-only

# Site: https://github.com/mugomes

import locale, os, sys, gettext

def getLang():
    def detectLanguage():
        lang, _ = locale.getlocale()
        if lang and lang != "C":
            return lang

        lang_env = os.environ.get("LANG", "")
        if lang_env:
            return lang_env.split(".")[0]  # remove encoding

        if sys.platform.startswith("win"):
            lang, _ = locale.getdefaultlocale()
            if lang:
                return lang

        return "en_US"

    lang = detectLanguage()
    try:
        langTrans = gettext.translation("messages", localedir='locales', languages=[lang])
        return langTrans.gettext
    except FileNotFoundError:
        return lambda s: s
