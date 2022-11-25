/*
&generator.GenDefinition{GenCommon:generator.GenCommon{Copyright:"", TargetImportPath:"github.com/zededa/terraform-provider"}, GenSchema:generator.GenSchema{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:true, IsNullable:false, IsStream:false, IsEmptyOmitted:true, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"SWState", Pkg:"models", PkgAlias:"", AliasedType:"string", SwaggerType:"string", SwaggerFormat:"", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}{"SW_STATE_UNSPECIFIED", "SW_STATE_INITIAL", "SW_STATE_DOWNLOAD_IN_PROGRESS", "SW_STATE_DOWNLOADED", "SW_STATE_VERIFIED", "SW_STATE_INSTALLED", "SW_STATE_BOOTING", "SW_STATE_RUNNING", "SW_STATE_HALTING", "SW_STATE_HALTED", "SW_STATE_REFRESHING", "SW_STATE_PURGING", "SW_STATE_RESOLVING_TAG", "SW_STATE_RESOLVED_TAG", "SW_STATE_CREATING_VOLUME", "SW_STATE_CREATED_VOLUME", "SW_STATE_VERIFYING", "SW_STATE_LOADING", "SW_STATE_LOADED", "SW_STATE_AWAITNETWORKINSTANCE"}}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:true, HasContextValidations:false, Required:false, HasSliceValidations:true, ItemsEnum:[]interface {}(nil)}, Example:"", OriginalName:"SWState", Name:"SWState", Suffix:"", Path:"", ValueExpression:"m", IndexVar:"i", KeyVar:"", Title:"State of EVE/Edge Application software", Description:"- SW_STATE_RESOLVING_TAG: Prior to DOWNLOAD_STARTED we go through these:Resolving an image tag - SW_STATE_RESOLVED_TAG: Tag has been resolved/failed - SW_STATE_CREATING_VOLUME: Prior to INSTALLED we go through these:Volume create in progress - SW_STATE_CREATED_VOLUME: Volume create done/failed - SW_STATE_VERIFYING: Verification in-progress - SW_STATE_LOADING: Loading blob in CAS - SW_STATE_LOADED: Loaded blob in CAS - SW_STATE_AWAITNETWORKINSTANCE: Wait for network instance", Location:"body", ReceiverName:"m", Items:(*generator.GenSchema)(nil), AllowsAdditionalItems:false, HasAdditionalItems:false, AdditionalItems:(*generator.GenSchema)(nil), Object:(*generator.GenSchema)(nil), XMLName:"", CustomTag:"", Properties:generator.GenSchemaList(nil), AllOf:generator.GenSchemaList(nil), HasAdditionalProperties:false, IsAdditionalProperties:false, AdditionalProperties:(*generator.GenSchema)(nil), StrictAdditionalProperties:false, ReadOnly:false, IsVirtual:false, IsBaseType:false, HasBaseType:false, IsSubType:false, IsExported:true, DiscriminatorField:"", DiscriminatorValue:"", Discriminates:map[string]string(nil), Parents:[]string(nil), IncludeValidator:true, IncludeModel:true, Default:"SW_STATE_UNSPECIFIED", WantsMarshalBinary:false, StructTags:[]string(nil), ExtraImports:map[string]string{}, ExternalDocs:(*spec.ExternalDocumentation)(nil)}, Package:"models", Imports:map[string]string{}, DefaultImports:map[string]string{"errors":"github.com/go-openapi/errors", "runtime":"github.com/go-openapi/runtime", "swag":"github.com/go-openapi/swag", "validate":"github.com/go-openapi/validate"}, ExtraSchemas:generator.GenSchemaList(nil), DependsOn:[]string(nil), External:false}
*/
