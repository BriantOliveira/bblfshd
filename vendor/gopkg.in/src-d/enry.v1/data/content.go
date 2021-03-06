package data

// CODE GENERATED AUTOMATICALLY WITH gopkg.in/src-d/enry.v1/internal/code-generator
// THIS FILE SHOULD NOT BE EDITED BY HAND
// Extracted from github/linguist commit: 4cd558c37482e8d2c535d8107f2d11b49afbc5b5

import "gopkg.in/toqueteos/substring.v1"

type languageMatcher func([]byte) []string

var ContentMatchers = map[string]languageMatcher{
	".asc": func(i []byte) []string {
		if asc_PublicKey_Matcher_0.Match(string(i)) {
			return []string{"Public Key"}
		} else if asc_AsciiDoc_Matcher_0.Match(string(i)) {
			return []string{"AsciiDoc"}
		} else if asc_AGSScript_Matcher_0.Match(string(i)) {
			return []string{"AGS Script"}
		}

		return nil
	},
	".bb": func(i []byte) []string {
		if bb_BlitzBasic_Matcher_0.Match(string(i)) || bb_BlitzBasic_Matcher_1.Match(string(i)) {
			return []string{"BlitzBasic"}
		} else if bb_BitBake_Matcher_0.Match(string(i)) {
			return []string{"BitBake"}
		}

		return nil
	},
	".builds": func(i []byte) []string {
		if builds_XML_Matcher_0.Match(string(i)) {
			return []string{"XML"}
		}

		return []string{"Text"}
	},
	".ch": func(i []byte) []string {
		if ch_xBase_Matcher_0.Match(string(i)) {
			return []string{"xBase"}
		}

		return nil
	},
	".cl": func(i []byte) []string {
		if cl_CommonLisp_Matcher_0.Match(string(i)) {
			return []string{"Common Lisp"}
		} else if cl_Cool_Matcher_0.Match(string(i)) {
			return []string{"Cool"}
		} else if cl_OpenCL_Matcher_0.Match(string(i)) {
			return []string{"OpenCL"}
		}

		return nil
	},
	".cls": func(i []byte) []string {
		if cls_TeX_Matcher_0.Match(string(i)) {
			return []string{"TeX"}
		}

		return nil
	},
	".cs": func(i []byte) []string {
		if cs_Smalltalk_Matcher_0.Match(string(i)) {
			return []string{"Smalltalk"}
		} else if cs_CSharp_Matcher_0.Match(string(i)) || cs_CSharp_Matcher_1.Match(string(i)) {
			return []string{"C#"}
		}

		return nil
	},
	".d": func(i []byte) []string {
		if d_D_Matcher_0.Match(string(i)) {
			return []string{"D"}
		} else if d_DTrace_Matcher_0.Match(string(i)) {
			return []string{"DTrace"}
		} else if d_Makefile_Matcher_0.Match(string(i)) {
			return []string{"Makefile"}
		}

		return nil
	},
	".ecl": func(i []byte) []string {
		if ecl_ECLiPSe_Matcher_0.Match(string(i)) {
			return []string{"ECLiPSe"}
		} else if ecl_ECL_Matcher_0.Match(string(i)) {
			return []string{"ECL"}
		}

		return nil
	},
	".es": func(i []byte) []string {
		if es_Erlang_Matcher_0.Match(string(i)) {
			return []string{"Erlang"}
		}

		return nil
	},
	".f": func(i []byte) []string {
		if f_Forth_Matcher_0.Match(string(i)) {
			return []string{"Forth"}
		} else if f_FilebenchWML_Matcher_0.Match(string(i)) {
			return []string{"Filebench WML"}
		} else if f_Fortran_Matcher_0.Match(string(i)) {
			return []string{"Fortran"}
		}

		return nil
	},
	".for": func(i []byte) []string {
		if for_Forth_Matcher_0.Match(string(i)) {
			return []string{"Forth"}
		} else if for_Fortran_Matcher_0.Match(string(i)) {
			return []string{"Fortran"}
		}

		return nil
	},
	".fr": func(i []byte) []string {
		if fr_Forth_Matcher_0.Match(string(i)) {
			return []string{"Forth"}
		} else if fr_Frege_Matcher_0.Match(string(i)) {
			return []string{"Frege"}
		}

		return []string{"Text"}
	},
	".fs": func(i []byte) []string {
		if fs_Forth_Matcher_0.Match(string(i)) {
			return []string{"Forth"}
		} else if fs_FSharp_Matcher_0.Match(string(i)) {
			return []string{"F#"}
		} else if fs_GLSL_Matcher_0.Match(string(i)) {
			return []string{"GLSL"}
		} else if fs_Filterscript_Matcher_0.Match(string(i)) {
			return []string{"Filterscript"}
		}

		return nil
	},
	".gs": func(i []byte) []string {
		if gs_Gosu_Matcher_0.Match(string(i)) {
			return []string{"Gosu"}
		}

		return nil
	},
	".h": func(i []byte) []string {
		if h_ObjectiveDashC_Matcher_0.Match(string(i)) {
			return []string{"Objective-C"}
		} else if h_CPlusPlus_Matcher_0.Match(string(i)) || h_CPlusPlus_Matcher_1.Match(string(i)) || h_CPlusPlus_Matcher_2.Match(string(i)) || h_CPlusPlus_Matcher_3.Match(string(i)) || h_CPlusPlus_Matcher_4.Match(string(i)) || h_CPlusPlus_Matcher_5.Match(string(i)) || h_CPlusPlus_Matcher_6.Match(string(i)) {
			return []string{"C++"}
		}

		return nil
	},
	".inc": func(i []byte) []string {
		if inc_PHP_Matcher_0.Match(string(i)) {
			return []string{"PHP"}
		} else if inc_POVDashRaySDL_Matcher_0.Match(string(i)) {
			return []string{"POV-Ray SDL"}
		}

		return nil
	},
	".l": func(i []byte) []string {
		if l_CommonLisp_Matcher_0.Match(string(i)) {
			return []string{"Common Lisp"}
		} else if l_Lex_Matcher_0.Match(string(i)) {
			return []string{"Lex"}
		} else if l_Roff_Matcher_0.Match(string(i)) {
			return []string{"Roff"}
		} else if l_PicoLisp_Matcher_0.Match(string(i)) {
			return []string{"PicoLisp"}
		}

		return nil
	},
	".ls": func(i []byte) []string {
		if ls_LoomScript_Matcher_0.Match(string(i)) {
			return []string{"LoomScript"}
		}

		return []string{"LiveScript"}
	},
	".lsp": func(i []byte) []string {
		if lsp_CommonLisp_Matcher_0.Match(string(i)) {
			return []string{"Common Lisp"}
		} else if lsp_NewLisp_Matcher_0.Match(string(i)) {
			return []string{"NewLisp"}
		}

		return nil
	},
	".lisp": func(i []byte) []string {
		if lisp_CommonLisp_Matcher_0.Match(string(i)) {
			return []string{"Common Lisp"}
		} else if lisp_NewLisp_Matcher_0.Match(string(i)) {
			return []string{"NewLisp"}
		}

		return nil
	},
	".m": func(i []byte) []string {
		if m_ObjectiveDashC_Matcher_0.Match(string(i)) {
			return []string{"Objective-C"}
		} else if m_Mercury_Matcher_0.Match(string(i)) {
			return []string{"Mercury"}
		} else if m_MUF_Matcher_0.Match(string(i)) {
			return []string{"MUF"}
		} else if m_M_Matcher_0.Match(string(i)) {
			return []string{"M"}
		} else if m_Mathematica_Matcher_0.Match(string(i)) {
			return []string{"Mathematica"}
		} else if m_Matlab_Matcher_0.Match(string(i)) {
			return []string{"Matlab"}
		} else if m_Limbo_Matcher_0.Match(string(i)) {
			return []string{"Limbo"}
		}

		return nil
	},
	".md": func(i []byte) []string {
		if md_Markdown_Matcher_0.Match(string(i)) || md_Markdown_Matcher_1.Match(string(i)) {
			return []string{"Markdown"}
		} else if md_GCCMachineDescription_Matcher_0.Match(string(i)) {
			return []string{"GCC Machine Description"}
		}

		return []string{"Markdown"}
	},
	".ml": func(i []byte) []string {
		if ml_OCaml_Matcher_0.Match(string(i)) {
			return []string{"OCaml"}
		} else if ml_StandardML_Matcher_0.Match(string(i)) {
			return []string{"Standard ML"}
		}

		return nil
	},
	".mod": func(i []byte) []string {
		if mod_XML_Matcher_0.Match(string(i)) {
			return []string{"XML"}
		} else if mod_ModulaDash2_Matcher_0.Match(string(i)) || mod_ModulaDash2_Matcher_1.Match(string(i)) {
			return []string{"Modula-2"}
		}

		return []string{"Linux Kernel Module", "AMPL"}
	},
	".ms": func(i []byte) []string {
		if ms_Roff_Matcher_0.Match(string(i)) {
			return []string{"Roff"}
		}

		return []string{"MAXScript"}
	},
	".n": func(i []byte) []string {
		if n_Roff_Matcher_0.Match(string(i)) {
			return []string{"Roff"}
		} else if n_Nemerle_Matcher_0.Match(string(i)) {
			return []string{"Nemerle"}
		}

		return nil
	},
	".ncl": func(i []byte) []string {
		if ncl_Text_Matcher_0.Match(string(i)) {
			return []string{"Text"}
		}

		return nil
	},
	".nl": func(i []byte) []string {
		if nl_NL_Matcher_0.Match(string(i)) {
			return []string{"NL"}
		}

		return []string{"NewLisp"}
	},
	".php": func(i []byte) []string {
		if php_Hack_Matcher_0.Match(string(i)) {
			return []string{"Hack"}
		} else if php_PHP_Matcher_0.Match(string(i)) {
			return []string{"PHP"}
		}

		return nil
	},
	".pl": func(i []byte) []string {
		if pl_Prolog_Matcher_0.Match(string(i)) {
			return []string{"Prolog"}
		} else if pl_Perl_Matcher_0.Match(string(i)) {
			return []string{"Perl"}
		} else if pl_Perl6_Matcher_0.Match(string(i)) {
			return []string{"Perl 6"}
		}

		return nil
	},
	".pm": func(i []byte) []string {
		if pm_Perl6_Matcher_0.Match(string(i)) {
			return []string{"Perl 6"}
		} else if pm_Perl_Matcher_0.Match(string(i)) {
			return []string{"Perl"}
		} else if pm_XPM_Matcher_0.Match(string(i)) {
			return []string{"XPM"}
		}

		return nil
	},
	".pod": func(i []byte) []string {
		if pod_Pod_Matcher_0.Match(string(i)) {
			return []string{"Pod"}
		}

		return []string{"Perl"}
	},
	"Pod": func(i []byte) []string {
		if Pod_Pod_Matcher_0.Match(string(i)) {
			return []string{"Pod"}
		}

		return []string{"Perl"}
	},
	"Perl": func(i []byte) []string {
		if Perl_Pod_Matcher_0.Match(string(i)) {
			return []string{"Pod"}
		}

		return []string{"Perl"}
	},
	".pro": func(i []byte) []string {
		if pro_Prolog_Matcher_0.Match(string(i)) {
			return []string{"Prolog"}
		} else if pro_INI_Matcher_0.Match(string(i)) {
			return []string{"INI"}
		} else if pro_QMake_Matcher_0.Match(string(i)) && pro_QMake_Matcher_1.Match(string(i)) {
			return []string{"QMake"}
		} else if pro_IDL_Matcher_0.Match(string(i)) {
			return []string{"IDL"}
		}

		return nil
	},
	".props": func(i []byte) []string {
		if props_XML_Matcher_0.Match(string(i)) {
			return []string{"XML"}
		} else if props_INI_Matcher_0.Match(string(i)) {
			return []string{"INI"}
		}

		return nil
	},
	".r": func(i []byte) []string {
		if r_Rebol_Matcher_0.Match(string(i)) {
			return []string{"Rebol"}
		} else if r_R_Matcher_0.Match(string(i)) {
			return []string{"R"}
		}

		return nil
	},
	".rno": func(i []byte) []string {
		if rno_RUNOFF_Matcher_0.Match(string(i)) {
			return []string{"RUNOFF"}
		} else if rno_Roff_Matcher_0.Match(string(i)) {
			return []string{"Roff"}
		}

		return nil
	},
	".rpy": func(i []byte) []string {
		if rpy_Python_Matcher_0.Match(string(i)) {
			return []string{"Python"}
		}

		return []string{"Ren'Py"}
	},
	".rs": func(i []byte) []string {
		if rs_Rust_Matcher_0.Match(string(i)) {
			return []string{"Rust"}
		} else if rs_RenderScript_Matcher_0.Match(string(i)) {
			return []string{"RenderScript"}
		}

		return nil
	},
	".sc": func(i []byte) []string {
		if sc_SuperCollider_Matcher_0.Match(string(i)) || sc_SuperCollider_Matcher_1.Match(string(i)) || sc_SuperCollider_Matcher_2.Match(string(i)) {
			return []string{"SuperCollider"}
		} else if sc_Scala_Matcher_0.Match(string(i)) || sc_Scala_Matcher_1.Match(string(i)) || sc_Scala_Matcher_2.Match(string(i)) {
			return []string{"Scala"}
		}

		return nil
	},
	".sql": func(i []byte) []string {
		if sql_PLpgSQL_Matcher_0.Match(string(i)) || sql_PLpgSQL_Matcher_1.Match(string(i)) || sql_PLpgSQL_Matcher_2.Match(string(i)) {
			return []string{"PLpgSQL"}
		} else if sql_SQLPL_Matcher_0.Match(string(i)) || sql_SQLPL_Matcher_1.Match(string(i)) {
			return []string{"SQLPL"}
		} else if sql_PLSQL_Matcher_0.Match(string(i)) || sql_PLSQL_Matcher_1.Match(string(i)) {
			return []string{"PLSQL"}
		} else if sql_SQL_Matcher_0.Match(string(i)) {
			return []string{"SQL"}
		}

		return nil
	},
	".srt": func(i []byte) []string {
		if srt_SubRipText_Matcher_0.Match(string(i)) {
			return []string{"SubRip Text"}
		}

		return nil
	},
	".t": func(i []byte) []string {
		if t_Turing_Matcher_0.Match(string(i)) {
			return []string{"Turing"}
		} else if t_Perl6_Matcher_0.Match(string(i)) {
			return []string{"Perl 6"}
		} else if t_Perl_Matcher_0.Match(string(i)) {
			return []string{"Perl"}
		}

		return nil
	},
	".toc": func(i []byte) []string {
		if toc_WorldofWarcraftAddonData_Matcher_0.Match(string(i)) {
			return []string{"World of Warcraft Addon Data"}
		} else if toc_TeX_Matcher_0.Match(string(i)) {
			return []string{"TeX"}
		}

		return nil
	},
	".ts": func(i []byte) []string {
		if ts_XML_Matcher_0.Match(string(i)) {
			return []string{"XML"}
		}

		return []string{"TypeScript"}
	},
	".tst": func(i []byte) []string {
		if tst_GAP_Matcher_0.Match(string(i)) {
			return []string{"GAP"}
		}

		return []string{"Scilab"}
	},
	".tsx": func(i []byte) []string {
		if tsx_TypeScript_Matcher_0.Match(string(i)) {
			return []string{"TypeScript"}
		} else if tsx_XML_Matcher_0.Match(string(i)) {
			return []string{"XML"}
		}

		return nil
	},
}

var (
	asc_PublicKey_Matcher_0                = substring.Regexp(`(?m)^(----[- ]BEGIN|ssh-(rsa|dss)) `)
	asc_AsciiDoc_Matcher_0                 = substring.Regexp(`(?m)^[=-]+(\s|\n)|{{[A-Za-z]`)
	asc_AGSScript_Matcher_0                = substring.Regexp(`(?m)^(\/\/.+|((import|export)\s+)?(function|int|float|char)\s+((room|repeatedly|on|game)_)?([A-Za-z]+[A-Za-z_0-9]+)\s*[;\(])`)
	bb_BlitzBasic_Matcher_0                = substring.Regexp(`(?m)^\s*; `)
	bb_BlitzBasic_Matcher_1                = substring.Regexp(`(?m)End Function`)
	bb_BitBake_Matcher_0                   = substring.Regexp(`(?m)^\s*(# |include|require)\b`)
	builds_XML_Matcher_0                   = substring.Regexp(`(?mi)^(\s*)(<Project|<Import|<Property|<?xml|xmlns)`)
	ch_xBase_Matcher_0                     = substring.Regexp(`(?mi)^\s*#\s*(if|ifdef|ifndef|define|command|xcommand|translate|xtranslate|include|pragma|undef)\b`)
	cl_CommonLisp_Matcher_0                = substring.Regexp(`(?mi)^\s*\((defun|in-package|defpackage) `)
	cl_Cool_Matcher_0                      = substring.Regexp(`(?m)^class`)
	cl_OpenCL_Matcher_0                    = substring.Regexp(`(?m)\/\* |\/\/ |^\}`)
	cls_TeX_Matcher_0                      = substring.Regexp(`(?m)\\\w+{`)
	cs_Smalltalk_Matcher_0                 = substring.Regexp(`(?m)![\w\s]+methodsFor: `)
	cs_CSharp_Matcher_0                    = substring.Regexp(`(?m)^\s*namespace\s*[\w\.]+\s*{`)
	cs_CSharp_Matcher_1                    = substring.Regexp(`(?m)^\s*\/\/`)
	d_D_Matcher_0                          = substring.Regexp(`(?m)^module\s+[\w.]*\s*;|import\s+[\w\s,.:]*;|\w+\s+\w+\s*\(.*\)(?:\(.*\))?\s*{[^}]*}|unittest\s*(?:\(.*\))?\s*{[^}]*}`)
	d_DTrace_Matcher_0                     = substring.Regexp(`(?m)^(\w+:\w*:\w*:\w*|BEGIN|END|provider\s+|(tick|profile)-\w+\s+{[^}]*}|#pragma\s+D\s+(option|attributes|depends_on)\s|#pragma\s+ident\s)`)
	d_Makefile_Matcher_0                   = substring.Regexp(`(?m)([\/\\].*:\s+.*\s\\$|: \\$|^ : |^[\w\s\/\\.]+\w+\.\w+\s*:\s+[\w\s\/\\.]+\w+\.\w+)`)
	ecl_ECLiPSe_Matcher_0                  = substring.Regexp(`(?m)^[^#]+:-`)
	ecl_ECL_Matcher_0                      = substring.Regexp(`(?m):=`)
	es_Erlang_Matcher_0                    = substring.Regexp(`(?m)^\s*(?:%%|main\s*\(.*?\)\s*->)`)
	f_Forth_Matcher_0                      = substring.Regexp(`(?m)^: `)
	f_FilebenchWML_Matcher_0               = substring.Regexp(`(?m)flowop`)
	f_Fortran_Matcher_0                    = substring.Regexp(`(?mi)^([c*][^abd-z]|      (subroutine|program|end|data)\s|\s*!)`)
	for_Forth_Matcher_0                    = substring.Regexp(`(?m)^: `)
	for_Fortran_Matcher_0                  = substring.Regexp(`(?mi)^([c*][^abd-z]|      (subroutine|program|end|data)\s|\s*!)`)
	fr_Forth_Matcher_0                     = substring.Regexp(`(?m)^(: |also |new-device|previous )`)
	fr_Frege_Matcher_0                     = substring.Regexp(`(?m)^\s*(import|module|package|data|type) `)
	fs_Forth_Matcher_0                     = substring.Regexp(`(?m)^(: |new-device)`)
	fs_FSharp_Matcher_0                    = substring.Regexp(`(?m)^\s*(#light|import|let|module|namespace|open|type)`)
	fs_GLSL_Matcher_0                      = substring.Regexp(`(?m)^\s*(#version|precision|uniform|varying|vec[234])`)
	fs_Filterscript_Matcher_0              = substring.Regexp(`(?m)#include|#pragma\s+(rs|version)|__attribute__`)
	gs_Gosu_Matcher_0                      = substring.Regexp(`(?m)^uses java\.`)
	h_ObjectiveDashC_Matcher_0             = substring.Regexp(`(?m)^\s*(@(interface|class|protocol|property|end|synchronised|selector|implementation)\b|#import\s+.+\.h[">])`)
	h_CPlusPlus_Matcher_0                  = substring.Regexp(`(?m)^\s*#\s*include <(cstdint|string|vector|map|list|array|bitset|queue|stack|forward_list|unordered_map|unordered_set|(i|o|io)stream)>`)
	h_CPlusPlus_Matcher_1                  = substring.Regexp(`(?m)^\s*template\s*<`)
	h_CPlusPlus_Matcher_2                  = substring.Regexp(`(?m)^[ \t]*try`)
	h_CPlusPlus_Matcher_3                  = substring.Regexp(`(?m)^[ \t]*catch\s*\(`)
	h_CPlusPlus_Matcher_4                  = substring.Regexp(`(?m)^[ \t]*(class|(using[ \t]+)?namespace)\s+\w+`)
	h_CPlusPlus_Matcher_5                  = substring.Regexp(`(?m)^[ \t]*(private|public|protected):$`)
	h_CPlusPlus_Matcher_6                  = substring.Regexp(`(?m)std::\w+`)
	inc_PHP_Matcher_0                      = substring.Regexp(`(?m)^<\?(?:php)?`)
	inc_POVDashRaySDL_Matcher_0            = substring.Regexp(`(?m)^\s*#(declare|local|macro|while)\s`)
	l_CommonLisp_Matcher_0                 = substring.Regexp(`(?m)\(def(un|macro)\s`)
	l_Lex_Matcher_0                        = substring.Regexp(`(?m)^(%[%{}]xs|<.*>)`)
	l_Roff_Matcher_0                       = substring.Regexp(`(?mi)^\.[a-z][a-z](\s|$)`)
	l_PicoLisp_Matcher_0                   = substring.Regexp(`(?m)^\((de|class|rel|code|data|must)\s`)
	ls_LoomScript_Matcher_0                = substring.Regexp(`(?m)^\s*package\s*[\w\.\/\*\s]*\s*{`)
	lsp_CommonLisp_Matcher_0               = substring.Regexp(`(?mi)^\s*\((defun|in-package|defpackage) `)
	lsp_NewLisp_Matcher_0                  = substring.Regexp(`(?m)^\s*\(define `)
	lisp_CommonLisp_Matcher_0              = substring.Regexp(`(?mi)^\s*\((defun|in-package|defpackage) `)
	lisp_NewLisp_Matcher_0                 = substring.Regexp(`(?m)^\s*\(define `)
	m_ObjectiveDashC_Matcher_0             = substring.Regexp(`(?m)^\s*(@(interface|class|protocol|property|end|synchronised|selector|implementation)\b|#import\s+.+\.h[">])`)
	m_Mercury_Matcher_0                    = substring.Regexp(`(?m):- module`)
	m_MUF_Matcher_0                        = substring.Regexp(`(?m)^: `)
	m_M_Matcher_0                          = substring.Regexp(`(?m)^\s*;`)
	m_Mathematica_Matcher_0                = substring.Regexp(`(?m)\*\)$`)
	m_Matlab_Matcher_0                     = substring.Regexp(`(?m)^\s*%`)
	m_Limbo_Matcher_0                      = substring.Regexp(`(?m)^\w+\s*:\s*module\s*{`)
	md_Markdown_Matcher_0                  = substring.Regexp(`(?mi)(^[-a-z0-9=#!\*\[|>])|<\/`)
	md_Markdown_Matcher_1                  = substring.Regexp(`^$`)
	md_GCCMachineDescription_Matcher_0     = substring.Regexp(`(?m)^(;;|\(define_)`)
	ml_OCaml_Matcher_0                     = substring.Regexp(`(?m)(^\s*module)|let rec |match\s+(\S+\s)+with`)
	ml_StandardML_Matcher_0                = substring.Regexp(`(?m)=> |case\s+(\S+\s)+of`)
	mod_XML_Matcher_0                      = substring.Regexp(`(?m)<!ENTITY `)
	mod_ModulaDash2_Matcher_0              = substring.Regexp(`(?mi)^\s*MODULE [\w\.]+;`)
	mod_ModulaDash2_Matcher_1              = substring.Regexp(`(?mi)^\s*END [\w\.]+;`)
	ms_Roff_Matcher_0                      = substring.Regexp(`(?mi)^[.'][a-z][a-z](\s|$)`)
	n_Roff_Matcher_0                       = substring.Regexp(`(?m)^[.']`)
	n_Nemerle_Matcher_0                    = substring.Regexp(`(?m)^(module|namespace|using)\s`)
	ncl_Text_Matcher_0                     = substring.Regexp(`(?m)THE_TITLE`)
	nl_NL_Matcher_0                        = substring.Regexp(`(?m)^(b|g)[0-9]+ `)
	php_Hack_Matcher_0                     = substring.Regexp(`(?m)<\?hh`)
	php_PHP_Matcher_0                      = substring.Regexp(`(?m)<?[^h]`)
	pl_Prolog_Matcher_0                    = substring.Regexp(`(?m)^[^#]*:-`)
	pl_Perl_Matcher_0                      = substring.Regexp(`(?m)use strict|use\s+v?5\.`)
	pl_Perl6_Matcher_0                     = substring.Regexp(`(?m)^(use v6|(my )?class|module)`)
	pm_Perl6_Matcher_0                     = substring.Regexp(`(?m)^\s*(?:use\s+v6\s*;|(?:\bmy\s+)?class|module)\b`)
	pm_Perl_Matcher_0                      = substring.Regexp(`(?m)\buse\s+(?:strict\b|v?5\.)`)
	pm_XPM_Matcher_0                       = substring.Regexp(`(?m)^\s*\/\* XPM \*\/`)
	pod_Pod_Matcher_0                      = substring.Regexp(`(?m)^=\w+\b`)
	Pod_Pod_Matcher_0                      = substring.Regexp(`(?m)^=\w+\b`)
	Perl_Pod_Matcher_0                     = substring.Regexp(`(?m)^=\w+\b`)
	pro_Prolog_Matcher_0                   = substring.Regexp(`(?m)^[^#]+:-`)
	pro_INI_Matcher_0                      = substring.Regexp(`(?m)last_client=`)
	pro_QMake_Matcher_0                    = substring.Regexp(`(?m)HEADERS`)
	pro_QMake_Matcher_1                    = substring.Regexp(`(?m)SOURCES`)
	pro_IDL_Matcher_0                      = substring.Regexp(`(?m)^\s*function[ \w,]+$`)
	props_XML_Matcher_0                    = substring.Regexp(`(?mi)^(\s*)(<Project|<Import|<Property|<?xml|xmlns)`)
	props_INI_Matcher_0                    = substring.Regexp(`(?mi)\w+\s*=\s*`)
	r_Rebol_Matcher_0                      = substring.Regexp(`(?mi)\bRebol\b`)
	r_R_Matcher_0                          = substring.Regexp(`(?m)<-|^\s*#`)
	rno_RUNOFF_Matcher_0                   = substring.Regexp(`(?mi)^\.!|^\.end lit(?:eral)?\b`)
	rno_Roff_Matcher_0                     = substring.Regexp(`(?m)^\.\\" `)
	rpy_Python_Matcher_0                   = substring.Regexp(`(?ms)(^(import|from|class|def)\s)`)
	rs_Rust_Matcher_0                      = substring.Regexp(`(?m)^(use |fn |mod |pub |macro_rules|impl|#!?\[)`)
	rs_RenderScript_Matcher_0              = substring.Regexp(`(?m)#include|#pragma\s+(rs|version)|__attribute__`)
	sc_SuperCollider_Matcher_0             = substring.Regexp(`(?m)\^(this|super)\.`)
	sc_SuperCollider_Matcher_1             = substring.Regexp(`(?m)^\s*(\+|\*)\s*\w+\s*{`)
	sc_SuperCollider_Matcher_2             = substring.Regexp(`(?m)^\s*~\w+\s*=\.`)
	sc_Scala_Matcher_0                     = substring.Regexp(`(?m)^\s*import (scala|java)\.`)
	sc_Scala_Matcher_1                     = substring.Regexp(`(?m)^\s*val\s+\w+\s*=`)
	sc_Scala_Matcher_2                     = substring.Regexp(`(?m)^\s*class\b`)
	sql_PLpgSQL_Matcher_0                  = substring.Regexp(`(?mi)^\\i\b|AS \$\$|LANGUAGE '?plpgsql'?`)
	sql_PLpgSQL_Matcher_1                  = substring.Regexp(`(?mi)SECURITY (DEFINER|INVOKER)`)
	sql_PLpgSQL_Matcher_2                  = substring.Regexp(`(?mi)BEGIN( WORK| TRANSACTION)?;`)
	sql_SQLPL_Matcher_0                    = substring.Regexp(`(?mi)(alter module)|(language sql)|(begin( NOT)+ atomic)`)
	sql_SQLPL_Matcher_1                    = substring.Regexp(`(?mi)signal SQLSTATE '[0-9]+'`)
	sql_PLSQL_Matcher_0                    = substring.Regexp(`(?mi)\$\$PLSQL_|XMLTYPE|sysdate|systimestamp|\.nextval|connect by|AUTHID (DEFINER|CURRENT_USER)`)
	sql_PLSQL_Matcher_1                    = substring.Regexp(`(?mi)constructor\W+function`)
	sql_SQL_Matcher_0                      = substring.Regexp(`(?mi)! /begin|boolean|package|exception`)
	srt_SubRipText_Matcher_0               = substring.Regexp(`(?m)^(\d{2}:\d{2}:\d{2},\d{3})\s*(-->)\s*(\d{2}:\d{2}:\d{2},\d{3})$`)
	t_Turing_Matcher_0                     = substring.Regexp(`(?m)^\s*%[ \t]+|^\s*var\s+\w+\s*:=\s*\w+`)
	t_Perl6_Matcher_0                      = substring.Regexp(`(?m)^\s*(?:use\s+v6\s*;|\bmodule\b|\b(?:my\s+)?class\b)`)
	t_Perl_Matcher_0                       = substring.Regexp(`(?m)\buse\s+(?:strict\b|v?5\.)`)
	toc_WorldofWarcraftAddonData_Matcher_0 = substring.Regexp(`(?m)^## |@no-lib-strip@`)
	toc_TeX_Matcher_0                      = substring.Regexp(`(?m)^\\(contentsline|defcounter|beamer|boolfalse)`)
	ts_XML_Matcher_0                       = substring.Regexp(`(?m)<TS`)
	tst_GAP_Matcher_0                      = substring.Regexp(`(?m)gap> `)
	tsx_TypeScript_Matcher_0               = substring.Regexp(`(?m)^\s*(import.+(from\s+|require\()['"]react|\/\/\/\s*<reference\s)`)
	tsx_XML_Matcher_0                      = substring.Regexp(`(?mi)^\s*<\?xml\s+version`)
)
