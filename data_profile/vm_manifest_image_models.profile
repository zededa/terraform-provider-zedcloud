/*
&generator.GenDefinition{GenCommon:generator.GenCommon{Copyright:"", TargetImportPath:"github.com/zededa/terraform-provider"}, GenSchema:generator.GenSchema{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:false, IsCustomFormatter:false, IsAliased:false, IsNullable:true, IsStream:false, IsEmptyOmitted:true, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:true, IsBaseType:false, HasDiscriminator:false, GoType:"VMManifestImage", Pkg:"models", PkgAlias:"", AliasedType:"", SwaggerType:"object", SwaggerFormat:"", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:true, HasContextValidations:true, Required:false, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, Example:"", OriginalName:"VMManifestImage", Name:"VMManifestImage", Suffix:"", Path:"", ValueExpression:"m", IndexVar:"i", KeyVar:"", Title:"", Description:"", Location:"body", ReceiverName:"m", Items:(*generator.GenSchema)(nil), AllowsAdditionalItems:false, HasAdditionalItems:false, AdditionalItems:(*generator.GenSchema)(nil), Object:(*generator.GenSchema)(nil), XMLName:"", CustomTag:"", Properties:generator.GenSchemaList{generator.GenSchema{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:true, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"bool", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"boolean", SwaggerFormat:"boolean", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:false, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, Example:"", OriginalName:"cleartext", Name:"cleartext", Suffix:"", Path:"\"cleartext\"", ValueExpression:"m.Cleartext", IndexVar:"i", KeyVar:"", Title:"", Description:"UI map: AppEditPage:DrivesPane:Cleartext, AppDetailsPage:DrivesPane:ClearText_Field", Location:"body", ReceiverName:"m", Items:(*generator.GenSchema)(nil), AllowsAdditionalItems:false, HasAdditionalItems:false, AdditionalItems:(*generator.GenSchema)(nil), Object:(*generator.GenSchema)(nil), XMLName:"", CustomTag:"", Properties:generator.GenSchemaList(nil), AllOf:generator.GenSchemaList(nil), HasAdditionalProperties:false, IsAdditionalProperties:false, AdditionalProperties:(*generator.GenSchema)(nil), StrictAdditionalProperties:false, ReadOnly:false, IsVirtual:false, IsBaseType:false, HasBaseType:false, IsSubType:false, IsExported:true, DiscriminatorField:"", DiscriminatorValue:"", Discriminates:map[string]string(nil), Parents:[]string(nil), IncludeValidator:true, IncludeModel:true, Default:interface {}(nil), WantsMarshalBinary:false, StructTags:[]string(nil), ExtraImports:map[string]string{}, ExternalDocs:(*spec.ExternalDocumentation)(nil)}, generator.GenSchema{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:true, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"string", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"string", SwaggerFormat:"", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:false, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, Example:"", OriginalName:"drvtype", Name:"drvtype", Suffix:"", Path:"\"drvtype\"", ValueExpression:"m.Drvtype", IndexVar:"i", KeyVar:"", Title:"enum: CDROM, HDD, NET", Description:"UI map: AppEditPage:DrivesPane:Drive_Type_Field, AppDetailsPage:DrivesPane:Drive_Type_Field", Location:"body", ReceiverName:"m", Items:(*generator.GenSchema)(nil), AllowsAdditionalItems:false, HasAdditionalItems:false, AdditionalItems:(*generator.GenSchema)(nil), Object:(*generator.GenSchema)(nil), XMLName:"", CustomTag:"", Properties:generator.GenSchemaList(nil), AllOf:generator.GenSchemaList(nil), HasAdditionalProperties:false, IsAdditionalProperties:false, AdditionalProperties:(*generator.GenSchema)(nil), StrictAdditionalProperties:false, ReadOnly:false, IsVirtual:false, IsBaseType:false, HasBaseType:false, IsSubType:false, IsExported:true, DiscriminatorField:"", DiscriminatorValue:"", Discriminates:map[string]string(nil), Parents:[]string(nil), IncludeValidator:true, IncludeModel:true, Default:interface {}(nil), WantsMarshalBinary:false, StructTags:[]string(nil), ExtraImports:map[string]string{}, ExternalDocs:(*spec.ExternalDocumentation)(nil)}, generator.GenSchema{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:true, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"bool", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"boolean", SwaggerFormat:"boolean", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:false, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, Example:"", OriginalName:"ignorepurge", Name:"ignorepurge", Suffix:"", Path:"\"ignorepurge\"", ValueExpression:"m.Ignorepurge", IndexVar:"i", KeyVar:"", Title:"", Description:"UI map: AppEditPage:DrivesPane:Ignorepurge, AppDetailsPage:DrivesPane:Ignorepurgee_Field", Location:"body", ReceiverName:"m", Items:(*generator.GenSchema)(nil), AllowsAdditionalItems:false, HasAdditionalItems:false, AdditionalItems:(*generator.GenSchema)(nil), Object:(*generator.GenSchema)(nil), XMLName:"", CustomTag:"", Properties:generator.GenSchemaList(nil), AllOf:generator.GenSchemaList(nil), HasAdditionalProperties:false, IsAdditionalProperties:false, AdditionalProperties:(*generator.GenSchema)(nil), StrictAdditionalProperties:false, ReadOnly:false, IsVirtual:false, IsBaseType:false, HasBaseType:false, IsSubType:false, IsExported:true, DiscriminatorField:"", DiscriminatorValue:"", Discriminates:map[string]string(nil), Parents:[]string(nil), IncludeValidator:true, IncludeModel:true, Default:interface {}(nil), WantsMarshalBinary:false, StructTags:[]string(nil), ExtraImports:map[string]string{}, ExternalDocs:(*spec.ExternalDocumentation)(nil)}, generator.GenSchema{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:true, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"string", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"string", SwaggerFormat:"", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:false, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, Example:"", OriginalName:"imagename", Name:"imagename", Suffix:"", Path:"\"imagename\"", ValueExpression:"m.Imagename", IndexVar:"i", KeyVar:"", Title:"", Description:"UI map: AppEditPage:DrivesPane:Image_Name_Field, AppDetailsPage:DrivesPane:Image_Name_Field", Location:"body", ReceiverName:"m", Items:(*generator.GenSchema)(nil), AllowsAdditionalItems:false, HasAdditionalItems:false, AdditionalItems:(*generator.GenSchema)(nil), Object:(*generator.GenSchema)(nil), XMLName:"", CustomTag:"", Properties:generator.GenSchemaList(nil), AllOf:generator.GenSchemaList(nil), HasAdditionalProperties:false, IsAdditionalProperties:false, AdditionalProperties:(*generator.GenSchema)(nil), StrictAdditionalProperties:false, ReadOnly:false, IsVirtual:false, IsBaseType:false, HasBaseType:false, IsSubType:false, IsExported:true, DiscriminatorField:"", DiscriminatorValue:"", Discriminates:map[string]string(nil), Parents:[]string(nil), IncludeValidator:true, IncludeModel:true, Default:interface {}(nil), WantsMarshalBinary:false, StructTags:[]string(nil), ExtraImports:map[string]string{}, ExternalDocs:(*spec.ExternalDocumentation)(nil)}, generator.GenSchema{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:true, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"string", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"string", SwaggerFormat:"uint64", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:false, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, Example:"", OriginalName:"maxsize", Name:"maxsize", Suffix:"", Path:"\"maxsize\"", ValueExpression:"m.Maxsize", IndexVar:"i", KeyVar:"", Title:"", Description:"UI map: AppEditPage:DrivesPane:Max_Size_Field, AppDetailsPage:DrivesPane:Max_Size_Field", Location:"body", ReceiverName:"m", Items:(*generator.GenSchema)(nil), AllowsAdditionalItems:false, HasAdditionalItems:false, AdditionalItems:(*generator.GenSchema)(nil), Object:(*generator.GenSchema)(nil), XMLName:"", CustomTag:"", Properties:generator.GenSchemaList(nil), AllOf:generator.GenSchemaList(nil), HasAdditionalProperties:false, IsAdditionalProperties:false, AdditionalProperties:(*generator.GenSchema)(nil), StrictAdditionalProperties:false, ReadOnly:false, IsVirtual:false, IsBaseType:false, HasBaseType:false, IsSubType:false, IsExported:true, DiscriminatorField:"", DiscriminatorValue:"", Discriminates:map[string]string(nil), Parents:[]string(nil), IncludeValidator:true, IncludeModel:true, Default:interface {}(nil), WantsMarshalBinary:false, StructTags:[]string(nil), ExtraImports:map[string]string{}, ExternalDocs:(*spec.ExternalDocumentation)(nil)}, generator.GenSchema{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:true, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"string", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"string", SwaggerFormat:"", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:false, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, Example:"", OriginalName:"mountpath", Name:"mountpath", Suffix:"", Path:"\"mountpath\"", ValueExpression:"m.Mountpath", IndexVar:"i", KeyVar:"", Title:"", Description:"UI map: AppEditPage:DrivesPane:Mountpath, AppDetailsPage:DrivesPane:Mountpath_Field", Location:"body", ReceiverName:"m", Items:(*generator.GenSchema)(nil), AllowsAdditionalItems:false, HasAdditionalItems:false, AdditionalItems:(*generator.GenSchema)(nil), Object:(*generator.GenSchema)(nil), XMLName:"", CustomTag:"", Properties:generator.GenSchemaList(nil), AllOf:generator.GenSchemaList(nil), HasAdditionalProperties:false, IsAdditionalProperties:false, AdditionalProperties:(*generator.GenSchema)(nil), StrictAdditionalProperties:false, ReadOnly:false, IsVirtual:false, IsBaseType:false, HasBaseType:false, IsSubType:false, IsExported:true, DiscriminatorField:"", DiscriminatorValue:"", Discriminates:map[string]string(nil), Parents:[]string(nil), IncludeValidator:true, IncludeModel:true, Default:interface {}(nil), WantsMarshalBinary:false, StructTags:[]string(nil), ExtraImports:map[string]string{}, ExternalDocs:(*spec.ExternalDocumentation)(nil)}, generator.GenSchema{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:true, IsMap:false, IsInterface:false, IsPrimitive:false, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:false, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"[]*Param", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"array", SwaggerFormat:"", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(0xc001150870), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:true, HasContextValidations:true, Required:false, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, Example:"", OriginalName:"params", Name:"params", Suffix:"", Path:"\"params\"", ValueExpression:"m.Params", IndexVar:"i", KeyVar:"", Title:"Not used by ZedUI", Description:"", Location:"body", ReceiverName:"m", Items:(*generator.GenSchema)(0xc0018f3c00), AllowsAdditionalItems:false, HasAdditionalItems:false, AdditionalItems:(*generator.GenSchema)(nil), Object:(*generator.GenSchema)(nil), XMLName:"", CustomTag:"", Properties:generator.GenSchemaList(nil), AllOf:generator.GenSchemaList(nil), HasAdditionalProperties:false, IsAdditionalProperties:false, AdditionalProperties:(*generator.GenSchema)(nil), StrictAdditionalProperties:false, ReadOnly:false, IsVirtual:false, IsBaseType:false, HasBaseType:false, IsSubType:false, IsExported:true, DiscriminatorField:"", DiscriminatorValue:"", Discriminates:map[string]string(nil), Parents:[]string(nil), IncludeValidator:true, IncludeModel:true, Default:interface {}(nil), WantsMarshalBinary:false, StructTags:[]string(nil), ExtraImports:map[string]string{}, ExternalDocs:(*spec.ExternalDocumentation)(nil)}, generator.GenSchema{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:true, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"bool", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"boolean", SwaggerFormat:"boolean", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:false, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, Example:"", OriginalName:"preserve", Name:"preserve", Suffix:"", Path:"\"preserve\"", ValueExpression:"m.Preserve", IndexVar:"i", KeyVar:"", Title:"", Description:"UI map: AppEditPage:DrivesPane:Preserve_Field, AppDetailsPage:DrivesPane:Preserve_Field", Location:"body", ReceiverName:"m", Items:(*generator.GenSchema)(nil), AllowsAdditionalItems:false, HasAdditionalItems:false, AdditionalItems:(*generator.GenSchema)(nil), Object:(*generator.GenSchema)(nil), XMLName:"", CustomTag:"", Properties:generator.GenSchemaList(nil), AllOf:generator.GenSchemaList(nil), HasAdditionalProperties:false, IsAdditionalProperties:false, AdditionalProperties:(*generator.GenSchema)(nil), StrictAdditionalProperties:false, ReadOnly:false, IsVirtual:false, IsBaseType:false, HasBaseType:false, IsSubType:false, IsExported:true, DiscriminatorField:"", DiscriminatorValue:"", Discriminates:map[string]string(nil), Parents:[]string(nil), IncludeValidator:true, IncludeModel:true, Default:interface {}(nil), WantsMarshalBinary:false, StructTags:[]string(nil), ExtraImports:map[string]string{}, ExternalDocs:(*spec.ExternalDocumentation)(nil)}, generator.GenSchema{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:true, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"bool", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"boolean", SwaggerFormat:"boolean", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:false, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, Example:"", OriginalName:"readonly", Name:"readonly", Suffix:"", Path:"\"readonly\"", ValueExpression:"m.Readonly", IndexVar:"i", KeyVar:"", Title:"Not used by ZedUI", Description:"", Location:"body", ReceiverName:"m", Items:(*generator.GenSchema)(nil), AllowsAdditionalItems:false, HasAdditionalItems:false, AdditionalItems:(*generator.GenSchema)(nil), Object:(*generator.GenSchema)(nil), XMLName:"", CustomTag:"", Properties:generator.GenSchemaList(nil), AllOf:generator.GenSchemaList(nil), HasAdditionalProperties:false, IsAdditionalProperties:false, AdditionalProperties:(*generator.GenSchema)(nil), StrictAdditionalProperties:false, ReadOnly:false, IsVirtual:false, IsBaseType:false, HasBaseType:false, IsSubType:false, IsExported:true, DiscriminatorField:"", DiscriminatorValue:"", Discriminates:map[string]string(nil), Parents:[]string(nil), IncludeValidator:true, IncludeModel:true, Default:interface {}(nil), WantsMarshalBinary:false, StructTags:[]string(nil), ExtraImports:map[string]string{}, ExternalDocs:(*spec.ExternalDocumentation)(nil)}, generator.GenSchema{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:true, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"string", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"string", SwaggerFormat:"", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:false, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, Example:"", OriginalName:"target", Name:"target", Suffix:"", Path:"\"target\"", ValueExpression:"m.Target", IndexVar:"i", KeyVar:"", Title:"enum: Disk, Kernel, Initrd, RamDisk", Description:"UI map: AppEditPage:DrivesPane:Target_Field, AppDetailsPage:DrivesPane:Target_Field", Location:"body", ReceiverName:"m", Items:(*generator.GenSchema)(nil), AllowsAdditionalItems:false, HasAdditionalItems:false, AdditionalItems:(*generator.GenSchema)(nil), Object:(*generator.GenSchema)(nil), XMLName:"", CustomTag:"", Properties:generator.GenSchemaList(nil), AllOf:generator.GenSchemaList(nil), HasAdditionalProperties:false, IsAdditionalProperties:false, AdditionalProperties:(*generator.GenSchema)(nil), StrictAdditionalProperties:false, ReadOnly:false, IsVirtual:false, IsBaseType:false, HasBaseType:false, IsSubType:false, IsExported:true, DiscriminatorField:"", DiscriminatorValue:"", Discriminates:map[string]string(nil), Parents:[]string(nil), IncludeValidator:true, IncludeModel:true, Default:interface {}(nil), WantsMarshalBinary:false, StructTags:[]string(nil), ExtraImports:map[string]string{}, ExternalDocs:(*spec.ExternalDocumentation)(nil)}, generator.GenSchema{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:true, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"string", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"string", SwaggerFormat:"", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:false, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, Example:"", OriginalName:"volumelabel", Name:"volumelabel", Suffix:"", Path:"\"volumelabel\"", ValueExpression:"m.Volumelabel", IndexVar:"i", KeyVar:"", Title:"", Description:"UI map: AppEditPage:DrivesPane:Volume_Label, AppDetailsPage:DrivesPane:Volume_Label", Location:"body", ReceiverName:"m", Items:(*generator.GenSchema)(nil), AllowsAdditionalItems:false, HasAdditionalItems:false, AdditionalItems:(*generator.GenSchema)(nil), Object:(*generator.GenSchema)(nil), XMLName:"", CustomTag:"", Properties:generator.GenSchemaList(nil), AllOf:generator.GenSchemaList(nil), HasAdditionalProperties:false, IsAdditionalProperties:false, AdditionalProperties:(*generator.GenSchema)(nil), StrictAdditionalProperties:false, ReadOnly:false, IsVirtual:false, IsBaseType:false, HasBaseType:false, IsSubType:false, IsExported:true, DiscriminatorField:"", DiscriminatorValue:"", Discriminates:map[string]string(nil), Parents:[]string(nil), IncludeValidator:true, IncludeModel:true, Default:interface {}(nil), WantsMarshalBinary:false, StructTags:[]string(nil), ExtraImports:map[string]string{}, ExternalDocs:(*spec.ExternalDocumentation)(nil)}}, AllOf:generator.GenSchemaList(nil), HasAdditionalProperties:false, IsAdditionalProperties:false, AdditionalProperties:(*generator.GenSchema)(nil), StrictAdditionalProperties:false, ReadOnly:false, IsVirtual:false, IsBaseType:false, HasBaseType:false, IsSubType:false, IsExported:true, DiscriminatorField:"", DiscriminatorValue:"", Discriminates:map[string]string(nil), Parents:[]string(nil), IncludeValidator:true, IncludeModel:true, Default:interface {}(nil), WantsMarshalBinary:true, StructTags:[]string(nil), ExtraImports:map[string]string{}, ExternalDocs:(*spec.ExternalDocumentation)(nil)}, Package:"models", Imports:map[string]string{}, DefaultImports:map[string]string{"errors":"github.com/go-openapi/errors", "runtime":"github.com/go-openapi/runtime", "swag":"github.com/go-openapi/swag", "validate":"github.com/go-openapi/validate"}, ExtraSchemas:generator.GenSchemaList(nil), DependsOn:[]string(nil), External:false}
*/
