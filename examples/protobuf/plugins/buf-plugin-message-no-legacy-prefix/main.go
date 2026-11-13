// buf-plugin-message-no-legacy-prefix is a buf check plugin that flags message names starting with "Legacy".
package main

import (
	"context"
	"strings"

	"buf.build/go/bufplugin/check"
	"buf.build/go/bufplugin/check/checkutil"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	check.Main(&check.Spec{
		Rules: []*check.RuleSpec{
			{
				ID:      "MESSAGE_NO_LEGACY_PREFIX",
				Type:    check.RuleTypeLint,
				Purpose: "Checks that message names do not start with Legacy.",
				Default: true,
				Handler: checkutil.NewMessageRuleHandler(checkMessage, checkutil.WithoutImports()),
			},
		},
	})
}

func checkMessage(_ context.Context, responseWriter check.ResponseWriter, _ check.Request, message protoreflect.MessageDescriptor) error {
	if strings.HasPrefix(string(message.Name()), "Legacy") {
		responseWriter.AddAnnotation(
			check.WithDescriptor(message),
			check.WithMessagef("Message name %q should not start with Legacy.", message.Name()),
		)
	}
	return nil
}
