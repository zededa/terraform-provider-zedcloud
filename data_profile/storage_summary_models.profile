/*
&generator.GenDefinition{GenCommon:generator.GenCommon{Copyright:"", TargetImportPath:"github.com/zededa/terraform-provider"}, GenSchema:generator.GenSchema{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:false, IsCustomFormatter:false, IsAliased:false, IsNullable:true, IsStream:false, IsEmptyOmitted:true, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:true, IsBaseType:false, HasDiscriminator:false, GoType:"StorageSummary", Pkg:"models", PkgAlias:"", AliasedType:"", SwaggerType:"object", SwaggerFormat:"", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:false, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, Example:"", OriginalName:"StorageSummary", Name:"StorageSummary", Suffix:"", Path:"", ValueExpression:"m", IndexVar:"i", KeyVar:"", Title:"Storage Summary", Description:"Storage Summary", Location:"body", ReceiverName:"m", Items:(*generator.GenSchema)(nil), AllowsAdditionalItems:false, HasAdditionalItems:false, AdditionalItems:(*generator.GenSchema)(nil), Object:(*generator.GenSchema)(nil), XMLName:"", CustomTag:"", Properties:generator.GenSchemaList{generator.GenSchema{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:true, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"float64", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"number", SwaggerFormat:"double", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:false, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, Example:"", OriginalName:"AllocatedMB", Name:"AllocatedMB", Suffix:"", Path:"\"AllocatedMB\"", ValueExpression:"m.AllocatedMB", IndexVar:"i", KeyVar:"", Title:"", Description:"Total reserved for running applications + temp. images etc", Location:"body", ReceiverName:"m", Items:(*generator.GenSchema)(nil), AllowsAdditionalItems:false, HasAdditionalItems:false, AdditionalItems:(*generator.GenSchema)(nil), Object:(*generator.GenSchema)(nil), XMLName:"", CustomTag:"", Properties:generator.GenSchemaList(nil), AllOf:generator.GenSchemaList(nil), HasAdditionalProperties:false, IsAdditionalProperties:false, AdditionalProperties:(*generator.GenSchema)(nil), StrictAdditionalProperties:false, ReadOnly:false, IsVirtual:false, IsBaseType:false, HasBaseType:false, IsSubType:false, IsExported:true, DiscriminatorField:"", DiscriminatorValue:"", Discriminates:map[string]string(nil), Parents:[]string(nil), IncludeValidator:true, IncludeModel:true, Default:interface {}(nil), WantsMarshalBinary:false, StructTags:[]string(nil), ExtraImports:map[string]string{}, ExternalDocs:(*spec.ExternalDocumentation)(nil)}, generator.GenSchema{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:true, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"float64", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"number", SwaggerFormat:"double", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:false, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, Example:"", OriginalName:"TotalMB", Name:"TotalMB", Suffix:"", Path:"\"TotalMB\"", ValueExpression:"m.TotalMB", IndexVar:"i", KeyVar:"", Title:"", Description:"Total Storage for the device in MBs", Location:"body", ReceiverName:"m", Items:(*generator.GenSchema)(nil), AllowsAdditionalItems:false, HasAdditionalItems:false, AdditionalItems:(*generator.GenSchema)(nil), Object:(*generator.GenSchema)(nil), XMLName:"", CustomTag:"", Properties:generator.GenSchemaList(nil), AllOf:generator.GenSchemaList(nil), HasAdditionalProperties:false, IsAdditionalProperties:false, AdditionalProperties:(*generator.GenSchema)(nil), StrictAdditionalProperties:false, ReadOnly:false, IsVirtual:false, IsBaseType:false, HasBaseType:false, IsSubType:false, IsExported:true, DiscriminatorField:"", DiscriminatorValue:"", Discriminates:map[string]string(nil), Parents:[]string(nil), IncludeValidator:true, IncludeModel:true, Default:interface {}(nil), WantsMarshalBinary:false, StructTags:[]string(nil), ExtraImports:map[string]string{}, ExternalDocs:(*spec.ExternalDocumentation)(nil)}, generator.GenSchema{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:true, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"float64", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"number", SwaggerFormat:"double", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:false, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, Example:"", OriginalName:"UsedMB", Name:"UsedMB", Suffix:"", Path:"\"UsedMB\"", ValueExpression:"m.UsedMB", IndexVar:"i", KeyVar:"", Title:"", Description:"How much is used within the allocated total storage", Location:"body", ReceiverName:"m", Items:(*generator.GenSchema)(nil), AllowsAdditionalItems:false, HasAdditionalItems:false, AdditionalItems:(*generator.GenSchema)(nil), Object:(*generator.GenSchema)(nil), XMLName:"", CustomTag:"", Properties:generator.GenSchemaList(nil), AllOf:generator.GenSchemaList(nil), HasAdditionalProperties:false, IsAdditionalProperties:false, AdditionalProperties:(*generator.GenSchema)(nil), StrictAdditionalProperties:false, ReadOnly:false, IsVirtual:false, IsBaseType:false, HasBaseType:false, IsSubType:false, IsExported:true, DiscriminatorField:"", DiscriminatorValue:"", Discriminates:map[string]string(nil), Parents:[]string(nil), IncludeValidator:true, IncludeModel:true, Default:interface {}(nil), WantsMarshalBinary:false, StructTags:[]string(nil), ExtraImports:map[string]string{}, ExternalDocs:(*spec.ExternalDocumentation)(nil)}}, AllOf:generator.GenSchemaList(nil), HasAdditionalProperties:false, IsAdditionalProperties:false, AdditionalProperties:(*generator.GenSchema)(nil), StrictAdditionalProperties:false, ReadOnly:false, IsVirtual:false, IsBaseType:false, HasBaseType:false, IsSubType:false, IsExported:true, DiscriminatorField:"", DiscriminatorValue:"", Discriminates:map[string]string(nil), Parents:[]string(nil), IncludeValidator:true, IncludeModel:true, Default:interface {}(nil), WantsMarshalBinary:true, StructTags:[]string(nil), ExtraImports:map[string]string{}, ExternalDocs:(*spec.ExternalDocumentation)(nil)}, Package:"models", Imports:map[string]string{}, DefaultImports:map[string]string{"errors":"github.com/go-openapi/errors", "runtime":"github.com/go-openapi/runtime", "swag":"github.com/go-openapi/swag", "validate":"github.com/go-openapi/validate"}, ExtraSchemas:generator.GenSchemaList(nil), DependsOn:[]string(nil), External:false}
*/
