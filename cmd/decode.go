// Copyright 2021 Adobe. All rights reserved.
// This file is licensed to you under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License. You may obtain a copy
// of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under
// the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR REPRESENTATIONS
// OF ANY KIND, either express or implied. See the License for the specific language
// governing permissions and limitations under the License.

package cmd

import (
	"fmt"

	"github.com/adobe/imscli/ims"
	"github.com/spf13/cobra"
)

func decodeCmd(imsConfig *ims.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "decode",
		Aliases: []string{"dec"},
		Short:   "Decode a JWT token.",
		Long:    "Decode a JWT token.",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true

			resp, err := imsConfig.DecodeToken()
			if err != nil {
				return fmt.Errorf("error decoding the token: %v", err)
			}
			for _, part := range resp {
				fmt.Println(part)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&imsConfig.Token, "token", "t", "", "Token.")

	return cmd
}
