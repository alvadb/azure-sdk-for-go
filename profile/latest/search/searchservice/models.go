package searchservice

import (
	 original "github.com/Azure/azure-sdk-for-go/service/search/2016-09-01/searchservice"
)

type (
	 IndexesClient = original.IndexesClient
	 CjkBigramTokenFilterScripts = original.CjkBigramTokenFilterScripts
	 EdgeNGramTokenFilterSide = original.EdgeNGramTokenFilterSide
	 IndexerExecutionStatus = original.IndexerExecutionStatus
	 IndexerStatus = original.IndexerStatus
	 MicrosoftStemmingTokenizerLanguage = original.MicrosoftStemmingTokenizerLanguage
	 MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguage
	 PhoneticEncoder = original.PhoneticEncoder
	 ScoringFunctionAggregation = original.ScoringFunctionAggregation
	 ScoringFunctionInterpolation = original.ScoringFunctionInterpolation
	 SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguage
	 StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguage
	 StopwordsList = original.StopwordsList
	 SuggesterSearchMode = original.SuggesterSearchMode
	 TokenCharacterKind = original.TokenCharacterKind
	 Analyzer = original.Analyzer
	 AnalyzeRequest = original.AnalyzeRequest
	 AnalyzeResult = original.AnalyzeResult
	 AnalyzerName = original.AnalyzerName
	 ASCIIFoldingTokenFilter = original.ASCIIFoldingTokenFilter
	 CharFilter = original.CharFilter
	 CharFilterName = original.CharFilterName
	 CjkBigramTokenFilter = original.CjkBigramTokenFilter
	 ClassicTokenizer = original.ClassicTokenizer
	 CommonGramTokenFilter = original.CommonGramTokenFilter
	 CorsOptions = original.CorsOptions
	 CustomAnalyzer = original.CustomAnalyzer
	 DataChangeDetectionPolicy = original.DataChangeDetectionPolicy
	 DataContainer = original.DataContainer
	 DataDeletionDetectionPolicy = original.DataDeletionDetectionPolicy
	 DataSource = original.DataSource
	 DataSourceCredentials = original.DataSourceCredentials
	 DataSourceListResult = original.DataSourceListResult
	 DataSourceType = original.DataSourceType
	 DataType = original.DataType
	 DictionaryDecompounderTokenFilter = original.DictionaryDecompounderTokenFilter
	 DistanceScoringFunction = original.DistanceScoringFunction
	 DistanceScoringParameters = original.DistanceScoringParameters
	 EdgeNGramTokenFilter = original.EdgeNGramTokenFilter
	 EdgeNGramTokenFilterV2 = original.EdgeNGramTokenFilterV2
	 EdgeNGramTokenizer = original.EdgeNGramTokenizer
	 ElisionTokenFilter = original.ElisionTokenFilter
	 Field = original.Field
	 FieldMapping = original.FieldMapping
	 FieldMappingFunction = original.FieldMappingFunction
	 FreshnessScoringFunction = original.FreshnessScoringFunction
	 FreshnessScoringParameters = original.FreshnessScoringParameters
	 HighWaterMarkChangeDetectionPolicy = original.HighWaterMarkChangeDetectionPolicy
	 Index = original.Index
	 Indexer = original.Indexer
	 IndexerExecutionInfo = original.IndexerExecutionInfo
	 IndexerExecutionResult = original.IndexerExecutionResult
	 IndexerListResult = original.IndexerListResult
	 IndexGetStatisticsResult = original.IndexGetStatisticsResult
	 IndexingParameters = original.IndexingParameters
	 IndexingSchedule = original.IndexingSchedule
	 IndexListResult = original.IndexListResult
	 ItemError = original.ItemError
	 KeepTokenFilter = original.KeepTokenFilter
	 KeywordMarkerTokenFilter = original.KeywordMarkerTokenFilter
	 KeywordTokenizer = original.KeywordTokenizer
	 KeywordTokenizerV2 = original.KeywordTokenizerV2
	 LengthTokenFilter = original.LengthTokenFilter
	 LimitTokenFilter = original.LimitTokenFilter
	 MagnitudeScoringFunction = original.MagnitudeScoringFunction
	 MagnitudeScoringParameters = original.MagnitudeScoringParameters
	 MappingCharFilter = original.MappingCharFilter
	 MicrosoftLanguageStemmingTokenizer = original.MicrosoftLanguageStemmingTokenizer
	 MicrosoftLanguageTokenizer = original.MicrosoftLanguageTokenizer
	 NGramTokenFilter = original.NGramTokenFilter
	 NGramTokenFilterV2 = original.NGramTokenFilterV2
	 NGramTokenizer = original.NGramTokenizer
	 PathHierarchyTokenizer = original.PathHierarchyTokenizer
	 PathHierarchyTokenizerV2 = original.PathHierarchyTokenizerV2
	 PatternAnalyzer = original.PatternAnalyzer
	 PatternCaptureTokenFilter = original.PatternCaptureTokenFilter
	 PatternReplaceCharFilter = original.PatternReplaceCharFilter
	 PatternReplaceTokenFilter = original.PatternReplaceTokenFilter
	 PatternTokenizer = original.PatternTokenizer
	 PhoneticTokenFilter = original.PhoneticTokenFilter
	 RegexFlags = original.RegexFlags
	 ScoringFunction = original.ScoringFunction
	 ScoringProfile = original.ScoringProfile
	 ShingleTokenFilter = original.ShingleTokenFilter
	 SnowballTokenFilter = original.SnowballTokenFilter
	 SoftDeleteColumnDeletionDetectionPolicy = original.SoftDeleteColumnDeletionDetectionPolicy
	 SQLIntegratedChangeTrackingPolicy = original.SQLIntegratedChangeTrackingPolicy
	 StandardAnalyzer = original.StandardAnalyzer
	 StandardTokenizer = original.StandardTokenizer
	 StandardTokenizerV2 = original.StandardTokenizerV2
	 StemmerOverrideTokenFilter = original.StemmerOverrideTokenFilter
	 StemmerTokenFilter = original.StemmerTokenFilter
	 StopAnalyzer = original.StopAnalyzer
	 StopwordsTokenFilter = original.StopwordsTokenFilter
	 Suggester = original.Suggester
	 SynonymTokenFilter = original.SynonymTokenFilter
	 TagScoringFunction = original.TagScoringFunction
	 TagScoringParameters = original.TagScoringParameters
	 TextWeights = original.TextWeights
	 TokenFilter = original.TokenFilter
	 TokenFilterName = original.TokenFilterName
	 TokenInfo = original.TokenInfo
	 Tokenizer = original.Tokenizer
	 TokenizerName = original.TokenizerName
	 TruncateTokenFilter = original.TruncateTokenFilter
	 UaxURLEmailTokenizer = original.UaxURLEmailTokenizer
	 UniqueTokenFilter = original.UniqueTokenFilter
	 WordDelimiterTokenFilter = original.WordDelimiterTokenFilter
	 ManagementClient = original.ManagementClient
	 DataSourcesClient = original.DataSourcesClient
	 IndexersClient = original.IndexersClient
)

const (
	 Han = original.Han
	 Hangul = original.Hangul
	 Hiragana = original.Hiragana
	 Katakana = original.Katakana
	 Back = original.Back
	 Front = original.Front
	 InProgress = original.InProgress
	 Reset = original.Reset
	 Success = original.Success
	 TransientFailure = original.TransientFailure
	 Error = original.Error
	 Running = original.Running
	 Unknown = original.Unknown
	 Arabic = original.Arabic
	 Bangla = original.Bangla
	 Bulgarian = original.Bulgarian
	 Catalan = original.Catalan
	 Croatian = original.Croatian
	 Czech = original.Czech
	 Danish = original.Danish
	 Dutch = original.Dutch
	 English = original.English
	 Estonian = original.Estonian
	 Finnish = original.Finnish
	 French = original.French
	 German = original.German
	 Greek = original.Greek
	 Gujarati = original.Gujarati
	 Hebrew = original.Hebrew
	 Hindi = original.Hindi
	 Hungarian = original.Hungarian
	 Icelandic = original.Icelandic
	 Indonesian = original.Indonesian
	 Italian = original.Italian
	 Kannada = original.Kannada
	 Latvian = original.Latvian
	 Lithuanian = original.Lithuanian
	 Malay = original.Malay
	 Malayalam = original.Malayalam
	 Marathi = original.Marathi
	 NorwegianBokmaal = original.NorwegianBokmaal
	 Polish = original.Polish
	 Portuguese = original.Portuguese
	 PortugueseBrazilian = original.PortugueseBrazilian
	 Punjabi = original.Punjabi
	 Romanian = original.Romanian
	 Russian = original.Russian
	 SerbianCyrillic = original.SerbianCyrillic
	 SerbianLatin = original.SerbianLatin
	 Slovak = original.Slovak
	 Slovenian = original.Slovenian
	 Spanish = original.Spanish
	 Swedish = original.Swedish
	 Tamil = original.Tamil
	 Telugu = original.Telugu
	 Turkish = original.Turkish
	 Ukrainian = original.Ukrainian
	 Urdu = original.Urdu
	 MicrosoftTokenizerLanguageBangla = original.MicrosoftTokenizerLanguageBangla
	 MicrosoftTokenizerLanguageBulgarian = original.MicrosoftTokenizerLanguageBulgarian
	 MicrosoftTokenizerLanguageCatalan = original.MicrosoftTokenizerLanguageCatalan
	 MicrosoftTokenizerLanguageChineseSimplified = original.MicrosoftTokenizerLanguageChineseSimplified
	 MicrosoftTokenizerLanguageChineseTraditional = original.MicrosoftTokenizerLanguageChineseTraditional
	 MicrosoftTokenizerLanguageCroatian = original.MicrosoftTokenizerLanguageCroatian
	 MicrosoftTokenizerLanguageCzech = original.MicrosoftTokenizerLanguageCzech
	 MicrosoftTokenizerLanguageDanish = original.MicrosoftTokenizerLanguageDanish
	 MicrosoftTokenizerLanguageDutch = original.MicrosoftTokenizerLanguageDutch
	 MicrosoftTokenizerLanguageEnglish = original.MicrosoftTokenizerLanguageEnglish
	 MicrosoftTokenizerLanguageFrench = original.MicrosoftTokenizerLanguageFrench
	 MicrosoftTokenizerLanguageGerman = original.MicrosoftTokenizerLanguageGerman
	 MicrosoftTokenizerLanguageGreek = original.MicrosoftTokenizerLanguageGreek
	 MicrosoftTokenizerLanguageGujarati = original.MicrosoftTokenizerLanguageGujarati
	 MicrosoftTokenizerLanguageHindi = original.MicrosoftTokenizerLanguageHindi
	 MicrosoftTokenizerLanguageIcelandic = original.MicrosoftTokenizerLanguageIcelandic
	 MicrosoftTokenizerLanguageIndonesian = original.MicrosoftTokenizerLanguageIndonesian
	 MicrosoftTokenizerLanguageItalian = original.MicrosoftTokenizerLanguageItalian
	 MicrosoftTokenizerLanguageJapanese = original.MicrosoftTokenizerLanguageJapanese
	 MicrosoftTokenizerLanguageKannada = original.MicrosoftTokenizerLanguageKannada
	 MicrosoftTokenizerLanguageKorean = original.MicrosoftTokenizerLanguageKorean
	 MicrosoftTokenizerLanguageMalay = original.MicrosoftTokenizerLanguageMalay
	 MicrosoftTokenizerLanguageMalayalam = original.MicrosoftTokenizerLanguageMalayalam
	 MicrosoftTokenizerLanguageMarathi = original.MicrosoftTokenizerLanguageMarathi
	 MicrosoftTokenizerLanguageNorwegianBokmaal = original.MicrosoftTokenizerLanguageNorwegianBokmaal
	 MicrosoftTokenizerLanguagePolish = original.MicrosoftTokenizerLanguagePolish
	 MicrosoftTokenizerLanguagePortuguese = original.MicrosoftTokenizerLanguagePortuguese
	 MicrosoftTokenizerLanguagePortugueseBrazilian = original.MicrosoftTokenizerLanguagePortugueseBrazilian
	 MicrosoftTokenizerLanguagePunjabi = original.MicrosoftTokenizerLanguagePunjabi
	 MicrosoftTokenizerLanguageRomanian = original.MicrosoftTokenizerLanguageRomanian
	 MicrosoftTokenizerLanguageRussian = original.MicrosoftTokenizerLanguageRussian
	 MicrosoftTokenizerLanguageSerbianCyrillic = original.MicrosoftTokenizerLanguageSerbianCyrillic
	 MicrosoftTokenizerLanguageSerbianLatin = original.MicrosoftTokenizerLanguageSerbianLatin
	 MicrosoftTokenizerLanguageSlovenian = original.MicrosoftTokenizerLanguageSlovenian
	 MicrosoftTokenizerLanguageSpanish = original.MicrosoftTokenizerLanguageSpanish
	 MicrosoftTokenizerLanguageSwedish = original.MicrosoftTokenizerLanguageSwedish
	 MicrosoftTokenizerLanguageTamil = original.MicrosoftTokenizerLanguageTamil
	 MicrosoftTokenizerLanguageTelugu = original.MicrosoftTokenizerLanguageTelugu
	 MicrosoftTokenizerLanguageThai = original.MicrosoftTokenizerLanguageThai
	 MicrosoftTokenizerLanguageUkrainian = original.MicrosoftTokenizerLanguageUkrainian
	 MicrosoftTokenizerLanguageUrdu = original.MicrosoftTokenizerLanguageUrdu
	 MicrosoftTokenizerLanguageVietnamese = original.MicrosoftTokenizerLanguageVietnamese
	 BeiderMorse = original.BeiderMorse
	 Caverphone1 = original.Caverphone1
	 Caverphone2 = original.Caverphone2
	 Cologne = original.Cologne
	 DoubleMetaphone = original.DoubleMetaphone
	 HaasePhonetik = original.HaasePhonetik
	 KoelnerPhonetik = original.KoelnerPhonetik
	 Metaphone = original.Metaphone
	 Nysiis = original.Nysiis
	 RefinedSoundex = original.RefinedSoundex
	 Soundex = original.Soundex
	 Average = original.Average
	 FirstMatching = original.FirstMatching
	 Maximum = original.Maximum
	 Minimum = original.Minimum
	 Sum = original.Sum
	 Constant = original.Constant
	 Linear = original.Linear
	 Logarithmic = original.Logarithmic
	 Quadratic = original.Quadratic
	 SnowballTokenFilterLanguageArmenian = original.SnowballTokenFilterLanguageArmenian
	 SnowballTokenFilterLanguageBasque = original.SnowballTokenFilterLanguageBasque
	 SnowballTokenFilterLanguageCatalan = original.SnowballTokenFilterLanguageCatalan
	 SnowballTokenFilterLanguageDanish = original.SnowballTokenFilterLanguageDanish
	 SnowballTokenFilterLanguageDutch = original.SnowballTokenFilterLanguageDutch
	 SnowballTokenFilterLanguageEnglish = original.SnowballTokenFilterLanguageEnglish
	 SnowballTokenFilterLanguageFinnish = original.SnowballTokenFilterLanguageFinnish
	 SnowballTokenFilterLanguageFrench = original.SnowballTokenFilterLanguageFrench
	 SnowballTokenFilterLanguageGerman = original.SnowballTokenFilterLanguageGerman
	 SnowballTokenFilterLanguageGerman2 = original.SnowballTokenFilterLanguageGerman2
	 SnowballTokenFilterLanguageHungarian = original.SnowballTokenFilterLanguageHungarian
	 SnowballTokenFilterLanguageItalian = original.SnowballTokenFilterLanguageItalian
	 SnowballTokenFilterLanguageKp = original.SnowballTokenFilterLanguageKp
	 SnowballTokenFilterLanguageLovins = original.SnowballTokenFilterLanguageLovins
	 SnowballTokenFilterLanguageNorwegian = original.SnowballTokenFilterLanguageNorwegian
	 SnowballTokenFilterLanguagePorter = original.SnowballTokenFilterLanguagePorter
	 SnowballTokenFilterLanguagePortuguese = original.SnowballTokenFilterLanguagePortuguese
	 SnowballTokenFilterLanguageRomanian = original.SnowballTokenFilterLanguageRomanian
	 SnowballTokenFilterLanguageRussian = original.SnowballTokenFilterLanguageRussian
	 SnowballTokenFilterLanguageSpanish = original.SnowballTokenFilterLanguageSpanish
	 SnowballTokenFilterLanguageSwedish = original.SnowballTokenFilterLanguageSwedish
	 SnowballTokenFilterLanguageTurkish = original.SnowballTokenFilterLanguageTurkish
	 StemmerTokenFilterLanguageArabic = original.StemmerTokenFilterLanguageArabic
	 StemmerTokenFilterLanguageArmenian = original.StemmerTokenFilterLanguageArmenian
	 StemmerTokenFilterLanguageBasque = original.StemmerTokenFilterLanguageBasque
	 StemmerTokenFilterLanguageBrazilian = original.StemmerTokenFilterLanguageBrazilian
	 StemmerTokenFilterLanguageBulgarian = original.StemmerTokenFilterLanguageBulgarian
	 StemmerTokenFilterLanguageCatalan = original.StemmerTokenFilterLanguageCatalan
	 StemmerTokenFilterLanguageCzech = original.StemmerTokenFilterLanguageCzech
	 StemmerTokenFilterLanguageDanish = original.StemmerTokenFilterLanguageDanish
	 StemmerTokenFilterLanguageDutch = original.StemmerTokenFilterLanguageDutch
	 StemmerTokenFilterLanguageDutchKp = original.StemmerTokenFilterLanguageDutchKp
	 StemmerTokenFilterLanguageEnglish = original.StemmerTokenFilterLanguageEnglish
	 StemmerTokenFilterLanguageFinnish = original.StemmerTokenFilterLanguageFinnish
	 StemmerTokenFilterLanguageFrench = original.StemmerTokenFilterLanguageFrench
	 StemmerTokenFilterLanguageGalician = original.StemmerTokenFilterLanguageGalician
	 StemmerTokenFilterLanguageGerman = original.StemmerTokenFilterLanguageGerman
	 StemmerTokenFilterLanguageGerman2 = original.StemmerTokenFilterLanguageGerman2
	 StemmerTokenFilterLanguageGreek = original.StemmerTokenFilterLanguageGreek
	 StemmerTokenFilterLanguageHindi = original.StemmerTokenFilterLanguageHindi
	 StemmerTokenFilterLanguageHungarian = original.StemmerTokenFilterLanguageHungarian
	 StemmerTokenFilterLanguageIndonesian = original.StemmerTokenFilterLanguageIndonesian
	 StemmerTokenFilterLanguageIrish = original.StemmerTokenFilterLanguageIrish
	 StemmerTokenFilterLanguageItalian = original.StemmerTokenFilterLanguageItalian
	 StemmerTokenFilterLanguageLatvian = original.StemmerTokenFilterLanguageLatvian
	 StemmerTokenFilterLanguageLightEnglish = original.StemmerTokenFilterLanguageLightEnglish
	 StemmerTokenFilterLanguageLightFinnish = original.StemmerTokenFilterLanguageLightFinnish
	 StemmerTokenFilterLanguageLightFrench = original.StemmerTokenFilterLanguageLightFrench
	 StemmerTokenFilterLanguageLightGerman = original.StemmerTokenFilterLanguageLightGerman
	 StemmerTokenFilterLanguageLightHungarian = original.StemmerTokenFilterLanguageLightHungarian
	 StemmerTokenFilterLanguageLightItalian = original.StemmerTokenFilterLanguageLightItalian
	 StemmerTokenFilterLanguageLightNorwegian = original.StemmerTokenFilterLanguageLightNorwegian
	 StemmerTokenFilterLanguageLightNynorsk = original.StemmerTokenFilterLanguageLightNynorsk
	 StemmerTokenFilterLanguageLightPortuguese = original.StemmerTokenFilterLanguageLightPortuguese
	 StemmerTokenFilterLanguageLightRussian = original.StemmerTokenFilterLanguageLightRussian
	 StemmerTokenFilterLanguageLightSpanish = original.StemmerTokenFilterLanguageLightSpanish
	 StemmerTokenFilterLanguageLightSwedish = original.StemmerTokenFilterLanguageLightSwedish
	 StemmerTokenFilterLanguageLovins = original.StemmerTokenFilterLanguageLovins
	 StemmerTokenFilterLanguageMinimalEnglish = original.StemmerTokenFilterLanguageMinimalEnglish
	 StemmerTokenFilterLanguageMinimalFrench = original.StemmerTokenFilterLanguageMinimalFrench
	 StemmerTokenFilterLanguageMinimalGalician = original.StemmerTokenFilterLanguageMinimalGalician
	 StemmerTokenFilterLanguageMinimalGerman = original.StemmerTokenFilterLanguageMinimalGerman
	 StemmerTokenFilterLanguageMinimalNorwegian = original.StemmerTokenFilterLanguageMinimalNorwegian
	 StemmerTokenFilterLanguageMinimalNynorsk = original.StemmerTokenFilterLanguageMinimalNynorsk
	 StemmerTokenFilterLanguageMinimalPortuguese = original.StemmerTokenFilterLanguageMinimalPortuguese
	 StemmerTokenFilterLanguageNorwegian = original.StemmerTokenFilterLanguageNorwegian
	 StemmerTokenFilterLanguagePorter2 = original.StemmerTokenFilterLanguagePorter2
	 StemmerTokenFilterLanguagePortuguese = original.StemmerTokenFilterLanguagePortuguese
	 StemmerTokenFilterLanguagePortugueseRslp = original.StemmerTokenFilterLanguagePortugueseRslp
	 StemmerTokenFilterLanguagePossessiveEnglish = original.StemmerTokenFilterLanguagePossessiveEnglish
	 StemmerTokenFilterLanguageRomanian = original.StemmerTokenFilterLanguageRomanian
	 StemmerTokenFilterLanguageRussian = original.StemmerTokenFilterLanguageRussian
	 StemmerTokenFilterLanguageSorani = original.StemmerTokenFilterLanguageSorani
	 StemmerTokenFilterLanguageSpanish = original.StemmerTokenFilterLanguageSpanish
	 StemmerTokenFilterLanguageSwedish = original.StemmerTokenFilterLanguageSwedish
	 StemmerTokenFilterLanguageTurkish = original.StemmerTokenFilterLanguageTurkish
	 StopwordsListArabic = original.StopwordsListArabic
	 StopwordsListArmenian = original.StopwordsListArmenian
	 StopwordsListBasque = original.StopwordsListBasque
	 StopwordsListBrazilian = original.StopwordsListBrazilian
	 StopwordsListBulgarian = original.StopwordsListBulgarian
	 StopwordsListCatalan = original.StopwordsListCatalan
	 StopwordsListCzech = original.StopwordsListCzech
	 StopwordsListDanish = original.StopwordsListDanish
	 StopwordsListDutch = original.StopwordsListDutch
	 StopwordsListEnglish = original.StopwordsListEnglish
	 StopwordsListFinnish = original.StopwordsListFinnish
	 StopwordsListFrench = original.StopwordsListFrench
	 StopwordsListGalician = original.StopwordsListGalician
	 StopwordsListGerman = original.StopwordsListGerman
	 StopwordsListGreek = original.StopwordsListGreek
	 StopwordsListHindi = original.StopwordsListHindi
	 StopwordsListHungarian = original.StopwordsListHungarian
	 StopwordsListIndonesian = original.StopwordsListIndonesian
	 StopwordsListIrish = original.StopwordsListIrish
	 StopwordsListItalian = original.StopwordsListItalian
	 StopwordsListLatvian = original.StopwordsListLatvian
	 StopwordsListNorwegian = original.StopwordsListNorwegian
	 StopwordsListPersian = original.StopwordsListPersian
	 StopwordsListPortuguese = original.StopwordsListPortuguese
	 StopwordsListRomanian = original.StopwordsListRomanian
	 StopwordsListRussian = original.StopwordsListRussian
	 StopwordsListSorani = original.StopwordsListSorani
	 StopwordsListSpanish = original.StopwordsListSpanish
	 StopwordsListSwedish = original.StopwordsListSwedish
	 StopwordsListThai = original.StopwordsListThai
	 StopwordsListTurkish = original.StopwordsListTurkish
	 AnalyzingInfixMatching = original.AnalyzingInfixMatching
	 Digit = original.Digit
	 Letter = original.Letter
	 Punctuation = original.Punctuation
	 Symbol = original.Symbol
	 Whitespace = original.Whitespace
	 DefaultBaseURI = original.DefaultBaseURI
)