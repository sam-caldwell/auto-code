package psqlToolChain

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest/contracts/common"
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
)

var Code = common.ToolChainCodeBlocks{
	Features: map[generic.Identifier]common.CodeBlock{
		"Enum": common.CodeBlock{
			/*
			 * Create all features needed to support a postgresql Enum type
			 */
			Forward: []string{
				`-- upsertEnum - creates an enumerated type with the given elements
				CREATE OR REPLACE FUNCTION upsertEnum(enum_name TEXT, elements TEXT[])
				RETURNS void AS $$
				BEGIN
				EXECUTE format('CREATE TYPE %I AS ENUM (%s)',
				enum_name,
				array_to_string(elements, ', '));
				END;
				$$ LANGUAGE plpgsql;`,
			},
			Rollback: []string{
				`DROP FUNCTION IF EXISTS upsertEnum`,
			},
		},
	},
}
