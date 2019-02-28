// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package storagedatalake

import original "github.com/Azure/azure-sdk-for-go/services/storage/datalake/2018-06-17/storagedatalake"

const (
	DefaultDNSSuffix = original.DefaultDNSSuffix
)

type BaseClient = original.BaseClient
type FilesystemClient = original.FilesystemClient
type PathGetPropertiesAction = original.PathGetPropertiesAction

const (
	GetAccessControl PathGetPropertiesAction = original.GetAccessControl
)

type PathLeaseAction = original.PathLeaseAction

const (
	Acquire PathLeaseAction = original.Acquire
	Break   PathLeaseAction = original.Break
	Change  PathLeaseAction = original.Change
	Release PathLeaseAction = original.Release
	Renew   PathLeaseAction = original.Renew
)

type PathRenameMode = original.PathRenameMode

const (
	Legacy PathRenameMode = original.Legacy
	Posix  PathRenameMode = original.Posix
)

type PathResourceType = original.PathResourceType

const (
	Directory PathResourceType = original.Directory
	File      PathResourceType = original.File
)

type PathUpdateAction = original.PathUpdateAction

const (
	Append           PathUpdateAction = original.Append
	Flush            PathUpdateAction = original.Flush
	SetAccessControl PathUpdateAction = original.SetAccessControl
	SetProperties    PathUpdateAction = original.SetProperties
)

type PathUpdateLeaseAction = original.PathUpdateLeaseAction

const (
	PathUpdateLeaseActionRelease PathUpdateLeaseAction = original.PathUpdateLeaseActionRelease
	PathUpdateLeaseActionRenew   PathUpdateLeaseAction = original.PathUpdateLeaseActionRenew
)

type DataLakeStorageError = original.DataLakeStorageError
type DataLakeStorageErrorError = original.DataLakeStorageErrorError
type Filesystem = original.Filesystem
type FilesystemList = original.FilesystemList
type Path = original.Path
type PathList = original.PathList
type ReadCloser = original.ReadCloser
type PathClient = original.PathClient

func New(xMsVersion string, accountName string) BaseClient {
	return original.New(xMsVersion, accountName)
}
func NewWithoutDefaults(xMsVersion string, accountName string, dNSSuffix string) BaseClient {
	return original.NewWithoutDefaults(xMsVersion, accountName, dNSSuffix)
}
func NewFilesystemClient(xMsVersion string, accountName string) FilesystemClient {
	return original.NewFilesystemClient(xMsVersion, accountName)
}
func PossiblePathGetPropertiesActionValues() []PathGetPropertiesAction {
	return original.PossiblePathGetPropertiesActionValues()
}
func PossiblePathLeaseActionValues() []PathLeaseAction {
	return original.PossiblePathLeaseActionValues()
}
func PossiblePathRenameModeValues() []PathRenameMode {
	return original.PossiblePathRenameModeValues()
}
func PossiblePathResourceTypeValues() []PathResourceType {
	return original.PossiblePathResourceTypeValues()
}
func PossiblePathUpdateActionValues() []PathUpdateAction {
	return original.PossiblePathUpdateActionValues()
}
func PossiblePathUpdateLeaseActionValues() []PathUpdateLeaseAction {
	return original.PossiblePathUpdateLeaseActionValues()
}
func NewPathClient(xMsVersion string, accountName string) PathClient {
	return original.NewPathClient(xMsVersion, accountName)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
