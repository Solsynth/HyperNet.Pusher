package provider

import (
	"fmt"

	"git.solsynth.dev/hypernet/nexus/pkg/nex/localize"
	"git.solsynth.dev/hypernet/pusher/pkg/pushkit"
)

func TranslateNotify(nty pushkit.Notification, lang string) pushkit.Notification {
	if nty.TranslateKey == nil {
		return nty
	}

	localizeKeys := map[string]string{
		"title":    fmt.Sprintf("%s.%s", *nty.TranslateKey, "subject"),
		"subtitle": fmt.Sprintf("%s.%s", *nty.TranslateKey, "subtitle"),
		"body":     fmt.Sprintf("%s.%s", *nty.TranslateKey, "body"),
	}

	for k, v := range localizeKeys {
		tmpl := localize.L.GetLocalizedString(v, lang)
		if args, ok := nty.TranslateArgs[k]; ok {
			anySlice := make([]any, len(args))
			for i, s := range args {
				anySlice[i] = s
			}
			str := fmt.Sprintf(tmpl, anySlice...)
			switch k {
			case "title":
				nty.Title = str
			case "subtitle":
				nty.Subtitle = str
			case "body":
				nty.Body = str
			}
		}

	}

	return nty
}
