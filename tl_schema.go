package mtproto

import (
	"fmt"
	"github.com/golang/glog"
	"runtime"
)

const (
	crc_boolFalse                                        = -1132882121
	crc_boolTrue                                         = -1720552011
	crc_error                                            = -994444869
	crc_null                                             = 1450380236
	crc_inputPeerEmpty                                   = 2134579434
	crc_inputPeerSelf                                    = 2107670217
	crc_inputPeerChat                                    = 396093539
	crc_inputUserEmpty                                   = -1182234929
	crc_inputUserSelf                                    = -138301121
	crc_inputPhoneContact                                = -208488460
	crc_inputFile                                        = -181407105
	crc_inputMediaEmpty                                  = -1771768449
	crc_inputMediaUploadedPhoto                          = 792191537
	crc_inputMediaPhoto                                  = -2114308294
	crc_inputMediaGeoPoint                               = -104578748
	crc_inputMediaContact                                = -1494984313
	crc_inputChatPhotoEmpty                              = 480546647
	crc_inputChatUploadedPhoto                           = -1837345356
	crc_inputChatPhoto                                   = -1991004873
	crc_inputGeoPointEmpty                               = -457104426
	crc_inputGeoPoint                                    = -206066487
	crc_inputPhotoEmpty                                  = 483901197
	crc_inputPhoto                                       = -74070332
	crc_inputFileLocation                                = 342061462
	crc_inputAppEvent                                    = 1996904104
	crc_peerUser                                         = -1649296275
	crc_peerChat                                         = -1160714821
	crc_storage_fileUnknown                              = -1432995067
	crc_storage_fileJpeg                                 = 8322574
	crc_storage_fileGif                                  = -891180321
	crc_storage_filePng                                  = 172975040
	crc_storage_fileMp3                                  = 1384777335
	crc_storage_fileMov                                  = 1258941372
	crc_storage_filePartial                              = 1086091090
	crc_storage_fileMp4                                  = -1278304028
	crc_storage_fileWebp                                 = 276907596
	crc_fileLocationUnavailable                          = 2086234950
	crc_fileLocation                                     = 1406570614
	crc_userEmpty                                        = 537022650
	crc_userProfilePhotoEmpty                            = 1326562017
	crc_userProfilePhoto                                 = -715532088
	crc_userStatusEmpty                                  = 164646985
	crc_userStatusOnline                                 = -306628279
	crc_userStatusOffline                                = 9203775
	crc_chatEmpty                                        = -1683826688
	crc_chat                                             = -652419756
	crc_chatForbidden                                    = 120753115
	crc_chatFull                                         = 771925524
	crc_chatParticipant                                  = -925415106
	crc_chatParticipantsForbidden                        = -57668565
	crc_chatParticipants                                 = 1061556205
	crc_chatPhotoEmpty                                   = 935395612
	crc_chatPhoto                                        = 1632839530
	crc_messageEmpty                                     = -2082087340
	crc_message                                          = 1157215293
	crc_messageService                                   = -1642487306
	crc_messageMediaEmpty                                = 1038967584
	crc_messageMediaPhoto                                = -1256047857
	crc_messageMediaGeo                                  = 1457575028
	crc_messageMediaContact                              = 1585262393
	crc_messageMediaUnsupported                          = -1618676578
	crc_messageActionEmpty                               = -1230047312
	crc_messageActionChatCreate                          = -1503425638
	crc_messageActionChatEditTitle                       = -1247687078
	crc_messageActionChatEditPhoto                       = 2144015272
	crc_messageActionChatDeletePhoto                     = -1780220945
	crc_messageActionChatAddUser                         = 1217033015
	crc_messageActionChatDeleteUser                      = -1297179892
	crc_dialog                                           = -455150117
	crc_photoEmpty                                       = 590459437
	crc_photo                                            = -1836524247
	crc_photoSizeEmpty                                   = 236446268
	crc_photoSize                                        = 2009052699
	crc_photoCachedSize                                  = -374917894
	crc_geoPointEmpty                                    = 286776671
	crc_geoPoint                                         = 541710092
	crc_auth_checkedPhone                                = -2128698738
	crc_auth_sentCode                                    = 1577067778
	crc_auth_authorization                               = -855308010
	crc_auth_exportedAuthorization                       = -543777747
	crc_inputNotifyPeer                                  = -1195615476
	crc_inputNotifyUsers                                 = 423314455
	crc_inputNotifyChats                                 = 1251338318
	crc_inputNotifyAll                                   = -1540769658
	crc_inputPeerNotifySettings                          = 949182130
	crc_peerNotifyEventsEmpty                            = -1378534221
	crc_peerNotifyEventsAll                              = 1830677896
	crc_peerNotifySettingsEmpty                          = 1889961234
	crc_peerNotifySettings                               = -1697798976
	crc_wallPaper                                        = -860866985
	crc_userFull                                         = 253890367
	crc_contact                                          = -116274796
	crc_importedContact                                  = -805141448
	crc_contactBlocked                                   = 1444661369
	crc_contactStatus                                    = -748155807
	crc_contacts_link                                    = 986597452
	crc_contacts_contacts                                = -353862078
	crc_contacts_contactsNotModified                     = -1219778094
	crc_contacts_importedContacts                        = 2010127419
	crc_contacts_blocked                                 = 471043349
	crc_contacts_blockedSlice                            = -1878523231
	crc_contacts_found                                   = 446822276
	crc_messages_dialogs                                 = 364538944
	crc_messages_dialogsSlice                            = 1910543603
	crc_messages_messages                                = -1938715001
	crc_messages_messagesSlice                           = 189033187
	crc_messages_chats                                   = 1694474197
	crc_messages_chatFull                                = -438840932
	crc_messages_affectedHistory                         = -1269012015
	crc_inputMessagesFilterEmpty                         = 1474492012
	crc_inputMessagesFilterPhotos                        = -1777752804
	crc_inputMessagesFilterVideo                         = -1614803355
	crc_inputMessagesFilterPhotoVideo                    = 1458172132
	crc_updateNewMessage                                 = 522914557
	crc_updateMessageID                                  = 1318109142
	crc_updateDeleteMessages                             = -1576161051
	crc_updateUserTyping                                 = 1548249383
	crc_updateChatUserTyping                             = -1704596961
	crc_updateChatParticipants                           = 125178264
	crc_updateUserStatus                                 = 469489699
	crc_updateUserName                                   = -1489818765
	crc_updateUserPhoto                                  = -1791935732
	crc_updateContactRegistered                          = 628472761
	crc_updateContactLink                                = -1657903163
	crc_updates_state                                    = -1519637954
	crc_updates_differenceEmpty                          = 1567990072
	crc_updates_difference                               = 16030880
	crc_updates_differenceSlice                          = -1459938943
	crc_updatesTooLong                                   = -484987010
	crc_updateShortMessage                               = -1857044719
	crc_updateShortChatMessage                           = 377562760
	crc_updateShort                                      = 2027216577
	crc_updatesCombined                                  = 1918567619
	crc_updates                                          = 1957577280
	crc_photos_photo                                     = 539045032
	crc_upload_file                                      = 157948117
	crc_dcOption                                         = 98092748
	crc_config                                           = -1669068444
	crc_nearestDc                                        = -1910892683
	crc_help_appUpdate                                   = -1987579119
	crc_help_noAppUpdate                                 = -1000708810
	crc_help_inviteText                                  = 415997816
	crc_inputPeerNotifyEventsEmpty                       = -265263912
	crc_inputPeerNotifyEventsAll                         = -395694988
	crc_photos_photos                                    = -1916114267
	crc_photos_photosSlice                               = 352657236
	crc_wallPaperSolid                                   = 1662091044
	crc_updateNewEncryptedMessage                        = 314359194
	crc_updateEncryptedChatTyping                        = 386986326
	crc_updateEncryption                                 = -1264392051
	crc_updateEncryptedMessagesRead                      = 956179895
	crc_encryptedChatEmpty                               = -1417756512
	crc_encryptedChatWaiting                             = 1006044124
	crc_encryptedChatRequested                           = -931638658
	crc_encryptedChat                                    = -94974410
	crc_encryptedChatDiscarded                           = 332848423
	crc_inputEncryptedChat                               = -247351839
	crc_encryptedFileEmpty                               = -1038136962
	crc_encryptedFile                                    = 1248893260
	crc_inputEncryptedFileEmpty                          = 406307684
	crc_inputEncryptedFileUploaded                       = 1690108678
	crc_inputEncryptedFile                               = 1511503333
	crc_inputEncryptedFileLocation                       = -182231723
	crc_encryptedMessage                                 = -317144808
	crc_encryptedMessageService                          = 594758406
	crc_messages_dhConfigNotModified                     = -1058912715
	crc_messages_dhConfig                                = 740433629
	crc_messages_sentEncryptedMessage                    = 1443858741
	crc_messages_sentEncryptedFile                       = -1802240206
	crc_inputFileBig                                     = -95482955
	crc_inputEncryptedFileBigUploaded                    = 767652808
	crc_storage_filePdf                                  = -1373745011
	crc_inputMessagesFilterDocument                      = -1629621880
	crc_inputMessagesFilterPhotoVideoDocuments           = -648121413
	crc_updateChatParticipantAdd                         = -364179876
	crc_updateChatParticipantDelete                      = 1851755554
	crc_updateDcOptions                                  = -1906403213
	crc_inputMediaUploadedDocument                       = -476700163
	crc_inputMediaDocument                               = 1523279502
	crc_messageMediaDocument                             = 2084836563
	crc_inputDocumentEmpty                               = 1928391342
	crc_inputDocument                                    = 410618194
	crc_inputDocumentFileLocation                        = 1125058340
	crc_documentEmpty                                    = 922273905
	crc_document                                         = -2027738169
	crc_help_support                                     = 398898678
	crc_notifyAll                                        = 1959820384
	crc_notifyChats                                      = -1073230141
	crc_notifyPeer                                       = -1613493288
	crc_notifyUsers                                      = -1261946036
	crc_updateUserBlocked                                = -2131957734
	crc_updateNotifySettings                             = -1094555409
	crc_sendMessageTypingAction                          = 381645902
	crc_sendMessageCancelAction                          = -44119819
	crc_sendMessageRecordVideoAction                     = -1584933265
	crc_sendMessageUploadVideoAction                     = -378127636
	crc_sendMessageRecordAudioAction                     = -718310409
	crc_sendMessageUploadAudioAction                     = -212740181
	crc_sendMessageUploadPhotoAction                     = -774682074
	crc_sendMessageUploadDocumentAction                  = -1441998364
	crc_sendMessageGeoLocationAction                     = 393186209
	crc_sendMessageChooseContactAction                   = 1653390447
	crc_updateServiceNotification                        = -337352679
	crc_userStatusRecently                               = -496024847
	crc_userStatusLastWeek                               = 129960444
	crc_userStatusLastMonth                              = 2011940674
	crc_updatePrivacy                                    = -298113238
	crc_inputPrivacyKeyStatusTimestamp                   = 1335282456
	crc_privacyKeyStatusTimestamp                        = -1137792208
	crc_inputPrivacyValueAllowContacts                   = 218751099
	crc_inputPrivacyValueAllowAll                        = 407582158
	crc_inputPrivacyValueAllowUsers                      = 320652927
	crc_inputPrivacyValueDisallowContacts                = 195371015
	crc_inputPrivacyValueDisallowAll                     = -697604407
	crc_inputPrivacyValueDisallowUsers                   = -1877932953
	crc_privacyValueAllowContacts                        = -123988
	crc_privacyValueAllowAll                             = 1698855810
	crc_privacyValueAllowUsers                           = 1297858060
	crc_privacyValueDisallowContacts                     = -125240806
	crc_privacyValueDisallowAll                          = -1955338397
	crc_privacyValueDisallowUsers                        = 209668535
	crc_account_privacyRules                             = 1430961007
	crc_accountDaysTTL                                   = -1194283041
	crc_updateUserPhone                                  = 314130811
	crc_disabledFeature                                  = -1369215196
	crc_documentAttributeImageSize                       = 1815593308
	crc_documentAttributeAnimated                        = 297109817
	crc_documentAttributeSticker                         = 1662637586
	crc_documentAttributeVideo                           = 250621158
	crc_documentAttributeAudio                           = -1739392570
	crc_documentAttributeFilename                        = 358154344
	crc_messages_stickersNotModified                     = -244016606
	crc_messages_stickers                                = -1970352846
	crc_stickerPack                                      = 313694676
	crc_messages_allStickersNotModified                  = -395967805
	crc_messages_allStickers                             = -302170017
	crc_account_noPassword                               = -1764049896
	crc_account_password                                 = 2081952796
	crc_updateReadHistoryInbox                           = -1721631396
	crc_updateReadHistoryOutbox                          = 791617983
	crc_messages_affectedMessages                        = -2066640507
	crc_contactLinkUnknown                               = 1599050311
	crc_contactLinkNone                                  = -17968211
	crc_contactLinkHasPhone                              = 646922073
	crc_contactLinkContact                               = -721239344
	crc_updateWebPage                                    = 2139689491
	crc_webPageEmpty                                     = -350980120
	crc_webPagePending                                   = -981018084
	crc_webPage                                          = 1594340540
	crc_messageMediaWebPage                              = -1557277184
	crc_authorization                                    = 2079516406
	crc_account_authorizations                           = 307276766
	crc_account_passwordSettings                         = -1212732749
	crc_account_passwordInputSettings                    = -2037289493
	crc_auth_passwordRecovery                            = 326715557
	crc_inputMediaVenue                                  = 673687578
	crc_messageMediaVenue                                = 2031269663
	crc_receivedNotifyMessage                            = -1551583367
	crc_chatInviteEmpty                                  = 1776236393
	crc_chatInviteExported                               = -64092740
	crc_chatInviteAlready                                = 1516793212
	crc_chatInvite                                       = -613092008
	crc_messageActionChatJoinedByLink                    = -123931160
	crc_updateReadMessagesContents                       = 1757493555
	crc_inputStickerSetEmpty                             = -4838507
	crc_inputStickerSetID                                = -1645763991
	crc_inputStickerSetShortName                         = -2044933984
	crc_stickerSet                                       = -852477119
	crc_messages_stickerSet                              = -1240849242
	crc_user                                             = 773059779
	crc_botCommand                                       = -1032140601
	crc_botInfo                                          = -1729618630
	crc_keyboardButton                                   = -1560655744
	crc_keyboardButtonRow                                = 2002815875
	crc_replyKeyboardHide                                = -1606526075
	crc_replyKeyboardForceReply                          = -200242528
	crc_replyKeyboardMarkup                              = 889353612
	crc_inputMessagesFilterUrl                           = 2129714567
	crc_inputPeerUser                                    = 2072935910
	crc_inputUser                                        = -668391402
	crc_messageEntityUnknown                             = -1148011883
	crc_messageEntityMention                             = -100378723
	crc_messageEntityHashtag                             = 1868782349
	crc_messageEntityBotCommand                          = 1827637959
	crc_messageEntityUrl                                 = 1859134776
	crc_messageEntityEmail                               = 1692693954
	crc_messageEntityBold                                = -1117713463
	crc_messageEntityItalic                              = -2106619040
	crc_messageEntityCode                                = 681706865
	crc_messageEntityPre                                 = 1938967520
	crc_messageEntityTextUrl                             = 1990644519
	crc_updateShortSentMessage                           = 301019932
	crc_inputPeerChannel                                 = 548253432
	crc_peerChannel                                      = -1109531342
	crc_channel                                          = 1158377749
	crc_channelForbidden                                 = 681420594
	crc_channelFull                                      = 1991201921
	crc_messageActionChannelCreate                       = -1781355374
	crc_messages_channelMessages                         = -1725551049
	crc_updateChannelTooLong                             = -352032773
	crc_updateChannel                                    = -1227598250
	crc_updateNewChannelMessage                          = 1656358105
	crc_updateReadChannelInbox                           = 1108669311
	crc_updateDeleteChannelMessages                      = -1015733815
	crc_updateChannelMessageViews                        = -1734268085
	crc_inputChannelEmpty                                = -292807034
	crc_inputChannel                                     = -1343524562
	crc_contacts_resolvedPeer                            = 2131196633
	crc_messageRange                                     = 182649427
	crc_updates_channelDifferenceEmpty                   = 1041346555
	crc_updates_channelDifferenceTooLong                 = 1788705589
	crc_updates_channelDifference                        = 543450958
	crc_channelMessagesFilterEmpty                       = -1798033689
	crc_channelMessagesFilter                            = -847783593
	crc_channelParticipant                               = 367766557
	crc_channelParticipantSelf                           = -1557620115
	crc_channelParticipantCreator                        = -471670279
	crc_channelParticipantsRecent                        = -566281095
	crc_channelParticipantsAdmins                        = -1268741783
	crc_channelParticipantsKicked                        = -1548400251
	crc_channels_channelParticipants                     = -177282392
	crc_channels_channelParticipant                      = -791039645
	crc_true                                             = 1072550713
	crc_chatParticipantCreator                           = -636267638
	crc_chatParticipantAdmin                             = -489233354
	crc_updateChatAdmins                                 = 1855224129
	crc_updateChatParticipantAdmin                       = -1232070311
	crc_messageActionChatMigrateTo                       = 1371385889
	crc_messageActionChannelMigrateFrom                  = -1336546578
	crc_channelParticipantsBots                          = -1328445861
	crc_inputReportReasonSpam                            = 1490799288
	crc_inputReportReasonViolence                        = 505595789
	crc_inputReportReasonPornography                     = 777640226
	crc_inputReportReasonOther                           = -512463606
	crc_updateNewStickerSet                              = 1753886890
	crc_updateStickerSetsOrder                           = 196268545
	crc_updateStickerSets                                = 1135492588
	crc_help_termsOfService                              = -236044656
	crc_foundGif                                         = 372165663
	crc_inputMediaGifExternal                            = 1212395773
	crc_messages_foundGifs                               = 1158290442
	crc_inputMessagesFilterGif                           = -3644025
	crc_updateSavedGifs                                  = -1821035490
	crc_updateBotInlineQuery                             = 1417832080
	crc_foundGifCached                                   = -1670052855
	crc_messages_savedGifsNotModified                    = -402498398
	crc_messages_savedGifs                               = 772213157
	crc_inputBotInlineMessageMediaAuto                   = 691006739
	crc_inputBotInlineMessageText                        = 1036876423
	crc_inputBotInlineResult                             = 750510426
	crc_botInlineMessageMediaAuto                        = 175419739
	crc_botInlineMessageText                             = -1937807902
	crc_botInlineResult                                  = -1679053127
	crc_messages_botResults                              = -858565059
	crc_inputMessagesFilterVoice                         = 1358283666
	crc_inputMessagesFilterMusic                         = 928101534
	crc_updateBotInlineSend                              = 239663460
	crc_inputPrivacyKeyChatInvite                        = -1107622874
	crc_privacyKeyChatInvite                             = 1343122938
	crc_updateEditChannelMessage                         = 457133559
	crc_exportedMessageLink                              = 524838915
	crc_messageFwdHeader                                 = -85986132
	crc_messageActionPinMessage                          = -1799538451
	crc_peerSettings                                     = -2122045747
	crc_updateChannelPinnedMessage                       = -1738988427
	crc_keyboardButtonUrl                                = 629866245
	crc_keyboardButtonCallback                           = 1748655686
	crc_keyboardButtonRequestPhone                       = -1318425559
	crc_keyboardButtonRequestGeoLocation                 = -59151553
	crc_auth_codeTypeSms                                 = 1923290508
	crc_auth_codeTypeCall                                = 1948046307
	crc_auth_codeTypeFlashCall                           = 577556219
	crc_auth_sentCodeTypeApp                             = 1035688326
	crc_auth_sentCodeTypeSms                             = -1073693790
	crc_auth_sentCodeTypeCall                            = 1398007207
	crc_auth_sentCodeTypeFlashCall                       = -1425815847
	crc_keyboardButtonSwitchInline                       = 90744648
	crc_replyInlineMarkup                                = 1218642516
	crc_messages_botCallbackAnswer                       = 911761060
	crc_updateBotCallbackQuery                           = -415938591
	crc_messages_messageEditData                         = 649453030
	crc_updateEditMessage                                = -469536605
	crc_inputBotInlineMessageMediaGeo                    = -190472735
	crc_inputBotInlineMessageMediaVenue                  = -1431327288
	crc_inputBotInlineMessageMediaContact                = 766443943
	crc_botInlineMessageMediaGeo                         = 982505656
	crc_botInlineMessageMediaVenue                       = 1130767150
	crc_botInlineMessageMediaContact                     = 904770772
	crc_inputBotInlineResultPhoto                        = -1462213465
	crc_inputBotInlineResultDocument                     = -459324
	crc_botInlineMediaResult                             = 400266251
	crc_inputBotInlineMessageID                          = -1995686519
	crc_updateInlineBotCallbackQuery                     = -103646630
	crc_inlineBotSwitchPM                                = 1008755359
	crc_messageEntityMentionName                         = 892193368
	crc_inputMessageEntityMentionName                    = 546203849
	crc_messages_peerDialogs                             = 863093588
	crc_topPeer                                          = -305282981
	crc_topPeerCategoryBotsPM                            = -1419371685
	crc_topPeerCategoryBotsInline                        = 344356834
	crc_topPeerCategoryCorrespondents                    = 104314861
	crc_topPeerCategoryGroups                            = -1122524854
	crc_topPeerCategoryChannels                          = 371037736
	crc_topPeerCategoryPeers                             = -75283823
	crc_contacts_topPeersNotModified                     = -567906571
	crc_contacts_topPeers                                = 1891070632
	crc_inputMessagesFilterChatPhotos                    = 975236280
	crc_updateReadChannelOutbox                          = 634833351
	crc_updateDraftMessage                               = -299124375
	crc_draftMessageEmpty                                = -1169445179
	crc_draftMessage                                     = -40996577
	crc_messageActionHistoryClear                        = -1615153660
	crc_updateReadFeaturedStickers                       = 1461528386
	crc_updateRecentStickers                             = -1706939360
	crc_messages_featuredStickersNotModified             = 82699215
	crc_messages_featuredStickers                        = -123893531
	crc_messages_recentStickersNotModified               = 186120336
	crc_messages_recentStickers                          = 1558317424
	crc_messages_archivedStickers                        = 1338747336
	crc_messages_stickerSetInstallResultSuccess          = 946083368
	crc_messages_stickerSetInstallResultArchive          = 904138920
	crc_stickerSetCovered                                = 1678812626
	crc_inputMediaPhotoExternal                          = 153267905
	crc_inputMediaDocumentExternal                       = -1225309387
	crc_updateConfig                                     = -1574314746
	crc_updatePtsChanged                                 = 861169551
	crc_messageActionGameScore                           = -1834538890
	crc_documentAttributeHasStickers                     = -1744710921
	crc_keyboardButtonGame                               = 1358175439
	crc_stickerSetMultiCovered                           = 872932635
	crc_maskCoords                                       = -1361650766
	crc_inputStickeredMediaPhoto                         = 1251549527
	crc_inputStickeredMediaDocument                      = 70813275
	crc_inputMediaGame                                   = -750828557
	crc_messageMediaGame                                 = -38694904
	crc_inputBotInlineMessageGame                        = 1262639204
	crc_inputBotInlineResultGame                         = 1336154098
	crc_game                                             = -1107729093
	crc_inputGameID                                      = 53231223
	crc_inputGameShortName                               = -1020139510
	crc_highScore                                        = 1493171408
	crc_messages_highScores                              = -1707344487
	crc_messages_chatsSlice                              = -1663561404
	crc_updateChannelWebPage                             = 1081547008
	crc_updates_differenceTooLong                        = 1258196845
	crc_sendMessageGamePlayAction                        = -580219064
	crc_webPageNotModified                               = -2054908813
	crc_textEmpty                                        = -599948721
	crc_textPlain                                        = 1950782688
	crc_textBold                                         = 1730456516
	crc_textItalic                                       = -653089380
	crc_textUnderline                                    = -1054465340
	crc_textStrike                                       = -1678197867
	crc_textFixed                                        = 1816074681
	crc_textUrl                                          = 1009288385
	crc_textEmail                                        = -564523562
	crc_textConcat                                       = 2120376535
	crc_pageBlockTitle                                   = 1890305021
	crc_pageBlockSubtitle                                = -1879401953
	crc_pageBlockAuthorDate                              = -1162877472
	crc_pageBlockHeader                                  = -1076861716
	crc_pageBlockSubheader                               = -248793375
	crc_pageBlockParagraph                               = 1182402406
	crc_pageBlockPreformatted                            = -1066346178
	crc_pageBlockFooter                                  = 1216809369
	crc_pageBlockDivider                                 = -618614392
	crc_pageBlockList                                    = 978896884
	crc_pageBlockBlockquote                              = 641563686
	crc_pageBlockPullquote                               = 1329878739
	crc_pageBlockPhoto                                   = -372860542
	crc_pageBlockVideo                                   = -640214938
	crc_pageBlockCover                                   = 972174080
	crc_pageBlockEmbed                                   = -840826671
	crc_pageBlockEmbedPost                               = 690781161
	crc_pageBlockSlideshow                               = 319588707
	crc_pagePart                                         = -1908433218
	crc_pageFull                                         = 1433323434
	crc_updatePhoneCall                                  = -1425052898
	crc_updateDialogPinned                               = -686710068
	crc_updatePinnedDialogs                              = -657787251
	crc_inputPrivacyKeyPhoneCall                         = -88417185
	crc_privacyKeyPhoneCall                              = 1030105979
	crc_pageBlockUnsupported                             = 324435594
	crc_pageBlockAnchor                                  = -837994576
	crc_pageBlockCollage                                 = 145955919
	crc_inputPhoneCall                                   = 506920429
	crc_phoneCallEmpty                                   = 1399245077
	crc_phoneCallWaiting                                 = 462375633
	crc_phoneCallRequested                               = -2089411356
	crc_phoneCall                                        = -1660057
	crc_phoneCallDiscarded                               = 1355435489
	crc_phoneConnection                                  = -1655957568
	crc_phoneCallProtocol                                = -1564789301
	crc_phone_phoneCall                                  = -326966976
	crc_phoneCallDiscardReasonMissed                     = -2048646399
	crc_phoneCallDiscardReasonDisconnect                 = -527056480
	crc_phoneCallDiscardReasonHangup                     = 1471006352
	crc_phoneCallDiscardReasonBusy                       = -84416311
	crc_inputMessagesFilterPhoneCalls                    = -2134272152
	crc_messageActionPhoneCall                           = -2132731265
	crc_invoice                                          = -1022713000
	crc_inputMediaInvoice                                = -1844103547
	crc_messageActionPaymentSentMe                       = -1892568281
	crc_messageMediaInvoice                              = -2074799289
	crc_keyboardButtonBuy                                = -1344716869
	crc_messageActionPaymentSent                         = 1080663248
	crc_payments_paymentForm                             = 1062645411
	crc_postAddress                                      = 512535275
	crc_paymentRequestedInfo                             = -1868808300
	crc_updateBotWebhookJSON                             = -2095595325
	crc_updateBotWebhookJSONQuery                        = -1684914010
	crc_updateBotShippingQuery                           = -523384512
	crc_updateBotPrecheckoutQuery                        = 1563376297
	crc_dataJSON                                         = 2104790276
	crc_labeledPrice                                     = -886477832
	crc_paymentCharge                                    = -368917890
	crc_paymentSavedCredentialsCard                      = -842892769
	crc_webDocument                                      = -971322408
	crc_inputWebDocument                                 = -1678949555
	crc_inputWebFileLocation                             = -1036396922
	crc_upload_webFile                                   = 568808380
	crc_payments_validatedRequestedInfo                  = -784000893
	crc_payments_paymentResult                           = 1314881805
	crc_payments_paymentVerficationNeeded                = 1800845601
	crc_payments_paymentReceipt                          = 1342771681
	crc_payments_savedInfo                               = -74456004
	crc_inputPaymentCredentialsSaved                     = -1056001329
	crc_inputPaymentCredentials                          = 873977640
	crc_account_tmpPassword                              = -614138572
	crc_shippingOption                                   = -1239335713
	crc_phoneCallAccepted                                = 1828732223
	crc_inputMessagesFilterRoundVoice                    = 2054952868
	crc_inputMessagesFilterRoundVideo                    = -1253451181
	crc_upload_fileCdnRedirect                           = -363659686
	crc_sendMessageRecordRoundAction                     = -1997373508
	crc_sendMessageUploadRoundAction                     = 608050278
	crc_upload_cdnFileReuploadNeeded                     = -290921362
	crc_upload_cdnFile                                   = -1449145777
	crc_cdnPublicKey                                     = -914167110
	crc_cdnConfig                                        = 1462101002
	crc_updateLangPackTooLong                            = 281165899
	crc_updateLangPack                                   = 1442983757
	crc_pageBlockChannel                                 = -283684427
	crc_inputStickerSetItem                              = -6249322
	crc_langPackString                                   = -892239370
	crc_langPackStringPluralized                         = 1816636575
	crc_langPackStringDeleted                            = 695856818
	crc_langPackDifference                               = -209337866
	crc_langPackLanguage                                 = 292985073
	crc_channelParticipantAdmin                          = -1473271656
	crc_channelParticipantBanned                         = 573315206
	crc_channelParticipantsBanned                        = 338142689
	crc_channelParticipantsSearch                        = 106343499
	crc_topPeerCategoryPhoneCalls                        = 511092620
	crc_pageBlockAudio                                   = 834148991
	crc_channelAdminRights                               = 1568467877
	crc_channelBannedRights                              = 1489977929
	crc_channelAdminLogEventActionChangeTitle            = -421545947
	crc_channelAdminLogEventActionChangeAbout            = 1427671598
	crc_channelAdminLogEventActionChangeUsername         = 1783299128
	crc_channelAdminLogEventActionChangePhoto            = -1204857405
	crc_channelAdminLogEventActionToggleInvites          = 460916654
	crc_channelAdminLogEventActionToggleSignatures       = 648939889
	crc_channelAdminLogEventActionUpdatePinned           = -370660328
	crc_channelAdminLogEventActionEditMessage            = 1889215493
	crc_channelAdminLogEventActionDeleteMessage          = 1121994683
	crc_channelAdminLogEventActionParticipantJoin        = 405815507
	crc_channelAdminLogEventActionParticipantLeave       = -124291086
	crc_channelAdminLogEventActionParticipantInvite      = -484690728
	crc_channelAdminLogEventActionParticipantToggleBan   = -422036098
	crc_channelAdminLogEventActionParticipantToggleAdmin = -714643696
	crc_channelAdminLogEvent                             = 995769920
	crc_channels_adminLogResults                         = -309659827
	crc_channelAdminLogEventsFilter                      = -368018716
	crc_messageActionScreenshotTaken                     = 1200788123
	crc_popularContact                                   = 1558266229
	crc_cdnFileHash                                      = 2012136335
	crc_inputMessagesFilterMyMentions                    = -1040652646
	crc_inputMessagesFilterMyMentionsUnread              = 1187706024
	crc_updateContactsReset                              = 1887741886
	crc_channelAdminLogEventActionChangeStickerSet       = -1312568665
	crc_updateFavedStickers                              = -451831443
	crc_messages_favedStickers                           = -209768682
	crc_messages_favedStickersNotModified                = -1634752813
	crc_updateChannelReadMessagesContents                = -1987495099
	crc_invokeAfterMsg                                   = -878758099
	crc_invokeAfterMsgs                                  = 1036301552
	crc_auth_checkPhone                                  = 1877286395
	crc_auth_sendCode                                    = -2035355412
	crc_auth_signUp                                      = 453408308
	crc_auth_signIn                                      = -1126886015
	crc_auth_logOut                                      = 1461180992
	crc_auth_resetAuthorizations                         = -1616179942
	crc_auth_sendInvites                                 = 1998331287
	crc_auth_exportAuthorization                         = -440401971
	crc_auth_importAuthorization                         = -470837741
	crc_account_registerDevice                           = 1669245048
	crc_account_unregisterDevice                         = 1707432768
	crc_account_updateNotifySettings                     = -2067899501
	crc_account_getNotifySettings                        = 313765169
	crc_account_resetNotifySettings                      = -612493497
	crc_account_updateProfile                            = 2018596725
	crc_account_updateStatus                             = 1713919532
	crc_account_getWallPapers                            = -1068696894
	crc_users_getUsers                                   = 227648840
	crc_users_getFullUser                                = -902781519
	crc_contacts_getStatuses                             = -995929106
	crc_contacts_getContacts                             = -1071414113
	crc_contacts_importContacts                          = 746589157
	crc_contacts_search                                  = 301470424
	crc_contacts_deleteContact                           = -1902823612
	crc_contacts_deleteContacts                          = 1504393374
	crc_contacts_block                                   = 858475004
	crc_contacts_unblock                                 = -448724803
	crc_contacts_getBlocked                              = -176409329
	crc_messages_getMessages                             = 1109588596
	crc_messages_getDialogs                              = 421243333
	crc_messages_getHistory                              = -1347868602
	crc_messages_search                                  = 60726944
	crc_messages_readHistory                             = 238054714
	crc_messages_deleteHistory                           = 469850889
	crc_messages_deleteMessages                          = -443640366
	crc_messages_receivedMessages                        = 94983360
	crc_messages_setTyping                               = -1551737264
	crc_messages_sendMessage                             = -91733382
	crc_messages_sendMedia                               = -923703407
	crc_messages_forwardMessages                         = 1888354709
	crc_messages_getChats                                = 1013621127
	crc_messages_getFullChat                             = 998448230
	crc_messages_editChatTitle                           = -599447467
	crc_messages_editChatPhoto                           = -900957736
	crc_messages_addChatUser                             = -106911223
	crc_messages_deleteChatUser                          = -530505962
	crc_messages_createChat                              = 164303470
	crc_updates_getState                                 = -304838614
	crc_updates_getDifference                            = 630429265
	crc_photos_updateProfilePhoto                        = -256159406
	crc_photos_uploadProfilePhoto                        = 1328726168
	crc_upload_saveFilePart                              = -1291540959
	crc_upload_getFile                                   = -475607115
	crc_help_getConfig                                   = -990308245
	crc_help_getNearestDc                                = 531836966
	crc_help_getAppUpdate                                = -1372724842
	crc_help_saveAppLog                                  = 1862465352
	crc_help_getInviteText                               = 1295590211
	crc_photos_deletePhotos                              = -2016444625
	crc_photos_getUserPhotos                             = -1848823128
	crc_messages_forwardMessage                          = 865483769
	crc_messages_getDhConfig                             = 651135312
	crc_messages_requestEncryption                       = -162681021
	crc_messages_acceptEncryption                        = 1035731989
	crc_messages_discardEncryption                       = -304536635
	crc_messages_setEncryptedTyping                      = 2031374829
	crc_messages_readEncryptedHistory                    = 2135648522
	crc_messages_sendEncrypted                           = -1451792525
	crc_messages_sendEncryptedFile                       = -1701831834
	crc_messages_sendEncryptedService                    = 852769188
	crc_messages_receivedQueue                           = 1436924774
	crc_upload_saveBigFilePart                           = -562337987
	crc_initConnection                                   = -951575130
	crc_help_getSupport                                  = -1663104819
	crc_auth_bindTempAuthKey                             = -841733627
	crc_contacts_exportCard                              = -2065352905
	crc_contacts_importCard                              = 1340184318
	crc_messages_readMessageContents                     = 916930423
	crc_account_checkUsername                            = 655677548
	crc_account_updateUsername                           = 1040964988
	crc_account_getPrivacy                               = -623130288
	crc_account_setPrivacy                               = -906486552
	crc_account_deleteAccount                            = 1099779595
	crc_account_getAccountTTL                            = 150761757
	crc_account_setAccountTTL                            = 608323678
	crc_invokeWithLayer                                  = -627372787
	crc_contacts_resolveUsername                         = -113456221
	crc_account_sendChangePhoneCode                      = 149257707
	crc_account_changePhone                              = 1891839707
	crc_messages_getAllStickers                          = 479598769
	crc_account_updateDeviceLocked                       = 954152242
	crc_account_getPassword                              = 1418342645
	crc_auth_checkPassword                               = 174260510
	crc_messages_getWebPagePreview                       = 623001124
	crc_account_getAuthorizations                        = -484392616
	crc_account_resetAuthorization                       = -545786948
	crc_account_getPasswordSettings                      = -1131605573
	crc_account_updatePasswordSettings                   = -92517498
	crc_auth_requestPasswordRecovery                     = -661144474
	crc_auth_recoverPassword                             = 1319464594
	crc_invokeWithoutUpdates                             = -1080796745
	crc_messages_exportChatInvite                        = 2106086025
	crc_messages_checkChatInvite                         = 1051570619
	crc_messages_importChatInvite                        = 1817183516
	crc_messages_getStickerSet                           = 639215886
	crc_messages_installStickerSet                       = -946871200
	crc_messages_uninstallStickerSet                     = -110209570
	crc_auth_importBotAuthorization                      = 1738800940
	crc_messages_startBot                                = -421563528
	crc_help_getAppChangelog                             = -1877938321
	crc_messages_reportSpam                              = -820669733
	crc_messages_getMessagesViews                        = -993483427
	crc_updates_getChannelDifference                     = 51854712
	crc_channels_readHistory                             = -871347913
	crc_channels_deleteMessages                          = -2067661490
	crc_channels_deleteUserHistory                       = -787622117
	crc_channels_reportSpam                              = -32999408
	crc_channels_getMessages                             = -1814580409
	crc_channels_getParticipants                         = 618237842
	crc_channels_getParticipant                          = 1416484774
	crc_channels_getChannels                             = 176122811
	crc_channels_getFullChannel                          = 141781513
	crc_channels_createChannel                           = -192332417
	crc_channels_editAbout                               = 333610782
	crc_channels_editAdmin                               = 548962836
	crc_channels_editTitle                               = 1450044624
	crc_channels_editPhoto                               = -248621111
	crc_channels_checkUsername                           = 283557164
	crc_channels_updateUsername                          = 890549214
	crc_channels_joinChannel                             = 615851205
	crc_channels_leaveChannel                            = -130635115
	crc_channels_inviteToChannel                         = 429865580
	crc_channels_exportInvite                            = -950663035
	crc_channels_deleteChannel                           = -1072619549
	crc_messages_toggleChatAdmins                        = -326379039
	crc_messages_editChatAdmin                           = -1444503762
	crc_messages_migrateChat                             = 363051235
	crc_messages_searchGlobal                            = -1640190800
	crc_account_reportPeer                               = -1374118561
	crc_messages_reorderStickerSets                      = 2016638777
	crc_help_getTermsOfService                           = 889286899
	crc_messages_getDocumentByHash                       = 864953444
	crc_messages_searchGifs                              = -1080395925
	crc_messages_getSavedGifs                            = -2084618926
	crc_messages_saveGif                                 = 846868683
	crc_messages_getInlineBotResults                     = 1364105629
	crc_messages_setInlineBotResults                     = -346119674
	crc_messages_sendInlineBotResult                     = -1318189314
	crc_channels_toggleInvites                           = 1231065863
	crc_channels_exportMessageLink                       = -934882771
	crc_channels_toggleSignatures                        = 527021574
	crc_messages_hideReportSpam                          = -1460572005
	crc_messages_getPeerSettings                         = 913498268
	crc_channels_updatePinnedMessage                     = -1490162350
	crc_auth_resendCode                                  = 1056025023
	crc_auth_cancelCode                                  = 520357240
	crc_messages_getMessageEditData                      = -39416522
	crc_messages_editMessage                             = -829299510
	crc_messages_editInlineBotMessage                    = 319564933
	crc_messages_getBotCallbackAnswer                    = -2130010132
	crc_messages_setBotCallbackAnswer                    = -712043766
	crc_contacts_getTopPeers                             = -728224331
	crc_contacts_resetTopPeerRating                      = 451113900
	crc_messages_getPeerDialogs                          = 764901049
	crc_messages_saveDraft                               = -1137057461
	crc_messages_getAllDrafts                            = 1782549861
	crc_account_sendConfirmPhoneCode                     = 353818557
	crc_account_confirmPhone                             = 1596029123
	crc_messages_getFeaturedStickers                     = 766298703
	crc_messages_readFeaturedStickers                    = 1527873830
	crc_messages_getRecentStickers                       = 1587647177
	crc_messages_saveRecentSticker                       = 958863608
	crc_messages_clearRecentStickers                     = -1986437075
	crc_messages_getArchivedStickers                     = 1475442322
	crc_channels_getAdminedPublicChannels                = -1920105769
	crc_auth_dropTempAuthKeys                            = -1907842680
	crc_messages_setGameScore                            = -1896289088
	crc_messages_setInlineGameScore                      = 363700068
	crc_messages_getMaskStickers                         = 1706608543
	crc_messages_getAttachedStickers                     = -866424884
	crc_messages_getGameHighScores                       = -400399203
	crc_messages_getInlineGameHighScores                 = 258170395
	crc_messages_getCommonChats                          = 218777796
	crc_messages_getAllChats                             = -341307408
	crc_help_setBotUpdatesStatus                         = -333262899
	crc_messages_getWebPage                              = 852135825
	crc_messages_toggleDialogPin                         = 847887978
	crc_messages_reorderPinnedDialogs                    = -1784678844
	crc_messages_getPinnedDialogs                        = -497756594
	crc_phone_requestCall                                = 1536537556
	crc_phone_acceptCall                                 = 1003664544
	crc_phone_discardCall                                = 2027164582
	crc_phone_receivedCall                               = 399855457
	crc_messages_reportEncryptedSpam                     = 1259113487
	crc_payments_getPaymentForm                          = -1712285883
	crc_payments_sendPaymentForm                         = 730364339
	crc_account_getTmpPassword                           = 1250046590
	crc_messages_setBotShippingResults                   = -436833542
	crc_messages_setBotPrecheckoutResults                = 163765653
	crc_upload_getWebFile                                = 619086221
	crc_bots_sendCustomRequest                           = -1440257555
	crc_bots_answerWebhookJSONQuery                      = -434028723
	crc_payments_getPaymentReceipt                       = -1601001088
	crc_payments_validateRequestedInfo                   = 1997180532
	crc_payments_getSavedInfo                            = 578650699
	crc_payments_clearSavedInfo                          = -667062079
	crc_phone_getCallConfig                              = 1430593449
	crc_phone_confirmCall                                = 788404002
	crc_phone_setCallRating                              = 475228724
	crc_phone_saveCallDebug                              = 662363518
	crc_upload_getCdnFile                                = 536919235
	crc_upload_reuploadCdnFile                           = 452533257
	crc_help_getCdnConfig                                = 1375900482
	crc_messages_uploadMedia                             = 1369162417
	crc_stickers_createStickerSet                        = -1680314774
	crc_langpack_getLangPack                             = -1699363442
	crc_langpack_getStrings                              = 773776152
	crc_langpack_getDifference                           = 187583869
	crc_langpack_getLanguages                            = -2146445955
	crc_channels_editBanned                              = -1076292147
	crc_channels_getAdminLog                             = 870184064
	crc_stickers_removeStickerFromSet                    = -143257775
	crc_stickers_changeStickerPosition                   = -4795190
	crc_stickers_addStickerToSet                         = -2041315650
	crc_messages_sendScreenshotNotification              = -914493408
	crc_upload_getCdnFileHashes                          = -149567365
	crc_messages_getUnreadMentions                       = 1180140658
	crc_messages_faveSticker                             = -1174420133
	crc_channels_setStickers                             = -359881479
	crc_contacts_resetSaved                              = -2020263951
	crc_messages_getFavedStickers                        = 567151374
	crc_channels_readMessageContents                     = -357180360
)

type TL_boolFalse struct {
}

type TL_boolTrue struct {
}

type TL_error struct {
	Code int32
	Text string
}

type TL_null struct {
}

type TL_inputPeerEmpty struct {
}

type TL_inputPeerSelf struct {
}

type TL_inputPeerChat struct {
	Chat_id int32
}

type TL_inputUserEmpty struct {
}

type TL_inputUserSelf struct {
}

type TL_inputPhoneContact struct {
	Client_id  int64
	Phone      string
	First_name string
	Last_name  string
}

type TL_inputFile struct {
	Id           int64
	Parts        int32
	Name         string
	Md5_checksum string
}

type TL_inputMediaEmpty struct {
}

type TL_inputMediaUploadedPhoto struct {
	Flags       int32
	File        TL // InputFile
	Caption     string
	Stickers    []TL // InputDocument
	Ttl_seconds int32
}

type TL_inputMediaPhoto struct {
	Flags       int32
	Id          TL // InputPhoto
	Caption     string
	Ttl_seconds int32
}

type TL_inputMediaGeoPoint struct {
	Geo_point TL // InputGeoPoint
}

type TL_inputMediaContact struct {
	Phone_number string
	First_name   string
	Last_name    string
}

type TL_inputChatPhotoEmpty struct {
}

type TL_inputChatUploadedPhoto struct {
	File TL // InputFile
}

type TL_inputChatPhoto struct {
	Id TL // InputPhoto
}

type TL_inputGeoPointEmpty struct {
}

type TL_inputGeoPoint struct {
	Lat  float64
	Long float64
}

type TL_inputPhotoEmpty struct {
}

type TL_inputPhoto struct {
	Id          int64
	Access_hash int64
}

type TL_inputFileLocation struct {
	Volume_id int64
	Local_id  int32
	Secret    int64
}

type TL_inputAppEvent struct {
	Time  float64
	_Type string
	Peer  int64
	Data  string
}

type TL_peerUser struct {
	User_id int32
}

type TL_peerChat struct {
	Chat_id int32
}

type TL_storage_fileUnknown struct {
}

type TL_storage_fileJpeg struct {
}

type TL_storage_fileGif struct {
}

type TL_storage_filePng struct {
}

type TL_storage_fileMp3 struct {
}

type TL_storage_fileMov struct {
}

type TL_storage_filePartial struct {
}

type TL_storage_fileMp4 struct {
}

type TL_storage_fileWebp struct {
}

type TL_fileLocationUnavailable struct {
	Volume_id int64
	Local_id  int32
	Secret    int64
}

type TL_fileLocation struct {
	Dc_id     int32
	Volume_id int64
	Local_id  int32
	Secret    int64
}

type TL_userEmpty struct {
	Id int32
}

type TL_userProfilePhotoEmpty struct {
}

type TL_userProfilePhoto struct {
	Photo_id    int64
	Photo_small TL // FileLocation
	Photo_big   TL // FileLocation
}

type TL_userStatusEmpty struct {
}

type TL_userStatusOnline struct {
	Expires int32
}

type TL_userStatusOffline struct {
	Was_online int32
}

type TL_chatEmpty struct {
	Id int32
}

type TL_chat struct {
	Flags int32
	// Creator	bool // flags_0?true
	// Kicked	bool // flags_1?true
	// Left	bool // flags_2?true
	// Admins_enabled	bool // flags_3?true
	// Admin	bool // flags_4?true
	// Deactivated	bool // flags_5?true
	Id                 int32
	Title              string
	Photo              TL // ChatPhoto
	Participants_count int32
	Date               int32
	Version            int32
	Migrated_to        TL // flags_6?InputChannel
}

type TL_chatForbidden struct {
	Id    int32
	Title string
}

type TL_chatFull struct {
	Id              int32
	Participants    TL   // ChatParticipants
	Chat_photo      TL   // Photo
	Notify_settings TL   // PeerNotifySettings
	Exported_invite TL   // ExportedChatInvite
	Bot_info        []TL // BotInfo
}

type TL_chatParticipant struct {
	User_id    int32
	Inviter_id int32
	Date       int32
}

type TL_chatParticipantsForbidden struct {
	Flags            int32
	Chat_id          int32
	Self_participant TL // flags_0?ChatParticipant
}

type TL_chatParticipants struct {
	Chat_id      int32
	Participants []TL // ChatParticipant
	Version      int32
}

type TL_chatPhotoEmpty struct {
}

type TL_chatPhoto struct {
	Photo_small TL // FileLocation
	Photo_big   TL // FileLocation
}

type TL_messageEmpty struct {
	Id int32
}

type TL_message struct {
	Flags           int32
	Id              int32
	Out             bool // flags_1?true
	Mentioned       bool // flags_4?true
	Media_unread    bool // flags_5?true
	Silent          bool // flags_13?true
	Post            bool // flags_14?true
	From_id         int32
	To_id           TL // Peer
	Fwd_from        TL // flags_2?MessageFwdHeader
	Via_bot_id      int32
	Reply_to_msg_id int32
	Date            int32
	Message         string
	Media           TL   // flags_9?MessageMedia
	Reply_markup    TL   // flags_6?ReplyMarkup
	Entities        []TL // MessageEntity
	Views           int32
	Edit_date       int32
	Post_author     string
	GroupedId       int64
	Action          TL
	Likes           int32
	Shares          int32
	Comments        int32
}

type TL_messageService struct {
	Flags int32
	// Out	bool // flags_1?true
	// Mentioned	bool // flags_4?true
	// Media_unread	bool // flags_5?true
	// Silent	bool // flags_13?true
	// Post	bool // flags_14?true
	Id              int32
	From_id         int32
	To_id           TL // Peer
	Reply_to_msg_id int32
	Date            int32
	Action          TL // MessageAction
}

type TL_messageMediaEmpty struct {
}

type TL_messageMediaPhoto struct {
	Flags       int32
	Photo       TL // flags_0?Photo
	Caption     string
	Ttl_seconds int32
}

type TL_messageMediaGeo struct {
	Geo TL // GeoPoint
}

type TL_messageMediaContact struct {
	Phone_number string
	First_name   string
	Last_name    string
	User_id      int32
}

type TL_messageMediaUnsupported struct {
}

type TL_messageActionEmpty struct {
}

type TL_messageActionChatCreate struct {
	Title string
	Users []int32
}

type TL_messageActionChatEditTitle struct {
	Title string
}

type TL_messageActionChatEditPhoto struct {
	Photo TL // Photo
}

type TL_messageActionChatDeletePhoto struct {
}

type TL_messageActionChatAddUser struct {
	Users []int32
}

type TL_messageActionChatDeleteUser struct {
	User_id int32
}

type TL_dialog struct {
	Flags int32
	// Pinned	bool // flags_2?true
	Peer                  TL // Peer
	Top_message           int32
	Read_inbox_max_id     int32
	Read_outbox_max_id    int32
	Unread_count          int32
	Unread_mentions_count int32
	Notify_settings       TL // PeerNotifySettings
	Pts                   int32
	Draft                 TL // flags_1?DraftMessage
}

type TL_photoEmpty struct {
	Id int64
}

type TL_photo struct {
	Flags int32
	// Has_stickers	bool // flags_0?true
	Id          int64
	Access_hash int64
	Date        int32
	Sizes       []TL // PhotoSize
}

type TL_photoSizeEmpty struct {
	_Type string
}

type TL_photoSize struct {
	_Type    string
	Location TL // FileLocation
	W        int32
	H        int32
	Size     int32
}

type TL_photoCachedSize struct {
	_Type    string
	Location TL // FileLocation
	W        int32
	H        int32
	Bytes    []byte
}

type TL_geoPointEmpty struct {
}

type TL_geoPoint struct {
	Long float64
	Lat  float64
}

type TL_auth_checkedPhone struct {
	Phone_registered TL // Bool
}

type TL_auth_sentCode struct {
	Flags           int32
	PhoneRegistered bool // flags_0?true
	Type            TL   // auth_SentCodeType
	PhoneCodeHash   string
	NextType        TL // flags_1?auth_CodeType
	Timeout         int32
	TermsOfService  TL
}

type TL_auth_authorization struct {
	Flags        int32
	Tmp_sessions int32
	User         TL // User
}

type TL_auth_exportedAuthorization struct {
	Id    int32
	Bytes []byte
}

type TL_inputNotifyPeer struct {
	Peer TL // InputPeer
}

type TL_inputNotifyUsers struct {
}

type TL_inputNotifyChats struct {
}

type TL_inputNotifyAll struct {
}

type TL_inputPeerNotifySettings struct {
	Flags int32
	// Show_previews	bool // flags_0?true
	// Silent	bool // flags_1?true
	Mute_until int32
	Sound      string
}

type TL_peerNotifyEventsEmpty struct {
}

type TL_peerNotifyEventsAll struct {
}

type TL_peerNotifySettingsEmpty struct {
}

type TL_peerNotifySettings struct {
	Flags int32
	// Show_previews	bool // flags_0?true
	// Silent	bool // flags_1?true
	Mute_until int32
	Sound      string
}

type TL_wallPaper struct {
	Id    int32
	Title string
	Sizes []TL // PhotoSize
	Color int32
}

type TL_userFull struct {
	Flags int32
	// Blocked	bool // flags_0?true
	// Phone_calls_available	bool // flags_4?true
	// Phone_calls_private	bool // flags_5?true
	User               TL // User
	About              string
	Link               TL // contacts_Link
	Profile_photo      TL // flags_2?Photo
	Notify_settings    TL // PeerNotifySettings
	Bot_info           TL // flags_3?BotInfo
	Common_chats_count int32
}

type TL_contact struct {
	User_id int32
	Mutual  TL // Bool
}

type TL_importedContact struct {
	User_id   int32
	Client_id int64
}

type TL_contactBlocked struct {
	User_id int32
	Date    int32
}

type TL_contactStatus struct {
	User_id int32
	Status  TL // UserStatus
}

type TL_contacts_link struct {
	My_link      TL // ContactLink
	Foreign_link TL // ContactLink
	User         TL // User
}

type TL_contacts_contacts struct {
	Contacts    []TL // Contact
	Saved_count int32
	Users       []TL // User
}


type TL_contacts_contactsNotModified struct {
}

type TL_contacts_importedContacts struct {
	Imported        []TL // ImportedContact
	Popular_invites []TL // PopularContact
	Retry_contacts  []int64
	Users           []TL // User
}

type TL_contacts_blocked struct {
	Blocked []TL // ContactBlocked
	Users   []TL // User
}

type TL_contacts_blockedSlice struct {
	Count   int32
	Blocked []TL // ContactBlocked
	Users   []TL // User
}

type TL_contacts_found struct {
	Results []TL // Peer
	Chats   []TL // Chat
	Users   []TL // User
}

type TL_messages_dialogs struct {
	Dialogs  []TL // Dialog
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TL_messages_dialogsSlice struct {
	Count    int32
	Dialogs  []TL // Dialog
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TL_messages_messages struct {
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TL_messages_messagesSlice struct {
	Count    int32
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TL_messages_chats struct {
	Chats []TL // Chat
}

type TL_messages_chatFull struct {
	Full_chat TL   // ChatFull
	Chats     []TL // Chat
	Users     []TL // User
}

type TL_messages_affectedHistory struct {
	Pts       int32
	Pts_count int32
	Offset    int32
}

type TL_inputMessagesFilterEmpty struct {
}

type TL_inputMessagesFilterPhotos struct {
}

type TL_inputMessagesFilterVideo struct {
}

type TL_inputMessagesFilterPhotoVideo struct {
}

type TL_updateNewMessage struct {
	Message   TL // Message
	Pts       int32
	Pts_count int32
}

type TL_updateMessageID struct {
	Id        int32
	Random_id int64
}

type TL_updateDeleteMessages struct {
	Messages  []int32
	Pts       int32
	Pts_count int32
}

type TL_updateUserTyping struct {
	User_id int32
	Action  TL // SendMessageAction
}

type TL_updateChatUserTyping struct {
	Chat_id int32
	User_id int32
	Action  TL // SendMessageAction
}

type TL_updateChatParticipants struct {
	Participants TL // ChatParticipants
}

type TL_updateUserStatus struct {
	User_id int32
	Status  TL // UserStatus
}

type TL_updateUserName struct {
	User_id    int32
	First_name string
	Last_name  string
	Username   string
}

type TL_updateUserPhoto struct {
	User_id  int32
	Date     int32
	Photo    TL // UserProfilePhoto
	Previous TL // Bool
}

type TL_updateContactRegistered struct {
	User_id int32
	Date    int32
}

type TL_updateContactLink struct {
	User_id      int32
	My_link      TL // ContactLink
	Foreign_link TL // ContactLink
}

type TL_updates_state struct {
	Pts          int32
	Qts          int32
	Date         int32
	Seq          int32
	Unread_count int32
}

type TL_updates_differenceEmpty struct {
	Date int32
	Seq  int32
}

type TL_updates_difference struct {
	New_messages           []TL // Message
	New_encrypted_messages []TL // EncryptedMessage
	Other_updates          []TL // Update
	Chats                  []TL // Chat
	Users                  []TL // User
	State                  TL   // updates_State
}

type TL_updates_differenceSlice struct {
	New_messages           []TL // Message
	New_encrypted_messages []TL // EncryptedMessage
	Other_updates          []TL // Update
	Chats                  []TL // Chat
	Users                  []TL // User
	Intermediate_state     TL   // updates_State
}

type TL_updatesTooLong struct {
}

type TL_updateShortMessage struct {
	Flags int32
	// Out	bool // flags_1?true
	// Mentioned	bool // flags_4?true
	// Media_unread	bool // flags_5?true
	// Silent	bool // flags_13?true
	Id              int32
	User_id         int32
	Message         string
	Pts             int32
	Pts_count       int32
	Date            int32
	Fwd_from        TL // flags_2?MessageFwdHeader
	Via_bot_id      int32
	Reply_to_msg_id int32
	Entities        []TL // MessageEntity
}

type TL_updateShortChatMessage struct {
	Flags int32
	// Out	bool // flags_1?true
	// Mentioned	bool // flags_4?true
	// Media_unread	bool // flags_5?true
	// Silent	bool // flags_13?true
	Id              int32
	From_id         int32
	Chat_id         int32
	Message         string
	Pts             int32
	Pts_count       int32
	Date            int32
	Fwd_from        TL // flags_2?MessageFwdHeader
	Via_bot_id      int32
	Reply_to_msg_id int32
	Entities        []TL // MessageEntity
}

type TL_updateShort struct {
	Update TL // Update
	Date   int32
}

type TL_updatesCombined struct {
	Updates   []TL // Update
	Users     []TL // User
	Chats     []TL // Chat
	Date      int32
	Seq_start int32
	Seq       int32
}

type TL_updates struct {
	Updates []TL // Update
	Users   []TL // User
	Chats   []TL // Chat
	Date    int32
	Seq     int32
}

type TL_photos_photo struct {
	Photo TL   // Photo
	Users []TL // User
}

type TL_upload_file struct {
	_Type TL // storage_FileType
	MTime int32
	Bytes []byte
}

type TL_dcOption struct {
	Flags     int32
	Ipv6      bool // flags_0?true
	MediaOnly bool // flags_1?true
	TCPoOnly  bool // flags_2?true
	Cdn       bool // flags_3?true
	Static    bool // flags_4?true
	Id        int32
	IpAddress string
	Port      int32
	Secret    []byte
}

type TL_config struct {
	Flags int32
	// Phonecalls_enabled	bool // flags_1?true
	Date                    int32
	Expires                 int32
	TestMode                TL // Bool
	ThisDC                  int32
	DcOptions               []TL // DcOption
	ChatSizeMax             int32
	MegagroupSizeMax        int32
	ForwardedCountMax       int32
	OnlineUpdatePeriodMs    int32
	OfflineBlurTimeoutMs    int32
	OfflineIdleTimeoutMs    int32
	OnlineCloudTimeoutMs    int32
	NotifyCloudDelayMs      int32
	NotifyDefaultDelayMs    int32
	ChatBigSize             int32
	PushChatPeriodMs        int32
	PushChatLimit           int32
	SavedGifsLimit          int32
	EditTimeLimit           int32
	RatingEDecay            int32
	StickersRecentLimit     int32
	StickersFavedLimit      int32
	ChannelsReadMediaPeriod int32
	TmpSessions             int32
	PinnedDialogsCountMax   int32
	CallReceiveTimeoutMs    int32
	CallRingTimeoutMs       int32
	CallConnectTimeoutMs    int32
	CallPacketTimeoutMs     int32
	MeUrlPrefix             string
	SuggestedLangCode       string
	LangPackVersion         int32
	Magic                   int32
	// Disabled_features        []TL // DisabledFeature
	CountDisableFeature int32
}

type TL_nearestDc struct {
	Country    string
	This_dc    int32
	Nearest_dc int32
}

type TL_help_appUpdate struct {
	Id       int32
	Critical TL // Bool
	Url      string
	Text     string
}

type TL_help_noAppUpdate struct {
}

type TL_help_inviteText struct {
	Message string
}

type TL_inputPeerNotifyEventsEmpty struct {
}

type TL_inputPeerNotifyEventsAll struct {
}

type TL_photos_photos struct {
	Photos []TL // Photo
	Users  []TL // User
}

type TL_photos_photosSlice struct {
	Count  int32
	Photos []TL // Photo
	Users  []TL // User
}

type TL_wallPaperSolid struct {
	Id       int32
	Title    string
	Bg_color int32
	Color    int32
}

type TL_updateNewEncryptedMessage struct {
	Message TL // EncryptedMessage
	Qts     int32
}

type TL_updateEncryptedChatTyping struct {
	Chat_id int32
}

type TL_updateEncryption struct {
	Chat TL // EncryptedChat
	Date int32
}

type TL_updateEncryptedMessagesRead struct {
	Chat_id  int32
	Max_date int32
	Date     int32
}

type TL_encryptedChatEmpty struct {
	Id int32
}

type TL_encryptedChatWaiting struct {
	Id             int32
	Access_hash    int64
	Date           int32
	Admin_id       int32
	Participant_id int32
}

type TL_encryptedChatRequested struct {
	Id             int32
	Access_hash    int64
	Date           int32
	Admin_id       int32
	Participant_id int32
	G_a            []byte
}

type TL_encryptedChat struct {
	Id              int32
	Access_hash     int64
	Date            int32
	Admin_id        int32
	Participant_id  int32
	G_a_or_b        []byte
	Key_fingerprint int64
}

type TL_encryptedChatDiscarded struct {
	Id int32
}

type TL_inputEncryptedChat struct {
	Chat_id     int32
	Access_hash int64
}

type TL_encryptedFileEmpty struct {
}

type TL_encryptedFile struct {
	Id              int64
	Access_hash     int64
	Size            int32
	Dc_id           int32
	Key_fingerprint int32
}

type TL_inputEncryptedFileEmpty struct {
}

type TL_inputEncryptedFileUploaded struct {
	Id              int64
	Parts           int32
	Md5_checksum    string
	Key_fingerprint int32
}

type TL_inputEncryptedFile struct {
	Id          int64
	Access_hash int64
}

type TL_inputEncryptedFileLocation struct {
	Id          int64
	Access_hash int64
}

type TL_encryptedMessage struct {
	Random_id int64
	Chat_id   int32
	Date      int32
	Bytes     []byte
	File      TL // EncryptedFile
}

type TL_encryptedMessageService struct {
	Random_id int64
	Chat_id   int32
	Date      int32
	Bytes     []byte
}

type TL_messages_dhConfigNotModified struct {
	Random []byte
}

type TL_messages_dhConfig struct {
	G       int32
	P       []byte
	Version int32
	Random  []byte
}

type TL_messages_sentEncryptedMessage struct {
	Date int32
}

type TL_messages_sentEncryptedFile struct {
	Date int32
	File TL // EncryptedFile
}

type TL_inputFileBig struct {
	Id    int64
	Parts int32
	Name  string
}

type TL_inputEncryptedFileBigUploaded struct {
	Id              int64
	Parts           int32
	Key_fingerprint int32
}

type TL_storage_filePdf struct {
}

type TL_inputMessagesFilterDocument struct {
}

type TL_inputMessagesFilterPhotoVideoDocuments struct {
}

type TL_updateChatParticipantAdd struct {
	Chat_id    int32
	User_id    int32
	Inviter_id int32
	Date       int32
	Version    int32
}

type TL_updateChatParticipantDelete struct {
	Chat_id int32
	User_id int32
	Version int32
}

type TL_updateDcOptions struct {
	Dc_options []TL // DcOption
}

type TL_inputMediaUploadedDocument struct {
	Flags       int32
	File        TL // InputFile
	Thumb       TL // flags_2?InputFile
	Mime_type   string
	Attributes  []TL // DocumentAttribute
	Caption     string
	Stickers    []TL // InputDocument
	Ttl_seconds int32
}

type TL_inputMediaDocument struct {
	Flags       int32
	Id          TL // InputDocument
	Caption     string
	Ttl_seconds int32
}

type TL_messageMediaDocument struct {
	Flags       int32
	Document    TL // flags_0?Document
	Caption     string
	Ttl_seconds int32
}

type TL_inputDocumentEmpty struct {
}

type TL_inputDocument struct {
	Id          int64
	Access_hash int64
}

type TL_inputDocumentFileLocation struct {
	Id          int64
	Access_hash int64
	Version     int32
}

type TL_documentEmpty struct {
	Id int64
}

type TL_document struct {
	Id          int64
	Access_hash int64
	Date        int32
	Mime_type   string
	Size        int32
	Thumb       TL // PhotoSize
	Dc_id       int32
	Version     int32
	Attributes  []TL // DocumentAttribute
}

type TL_help_support struct {
	Phone_number string
	User         TL // User
}

type TL_notifyAll struct {
}

type TL_notifyChats struct {
}

type TL_notifyPeer struct {
	Peer TL // Peer
}

type TL_notifyUsers struct {
}

type TL_updateUserBlocked struct {
	User_id int32
	Blocked TL // Bool
}

type TL_updateNotifySettings struct {
	Peer            TL // NotifyPeer
	Notify_settings TL // PeerNotifySettings
}

type TL_sendMessageTypingAction struct {
}

type TL_sendMessageCancelAction struct {
}

type TL_sendMessageRecordVideoAction struct {
}

type TL_sendMessageUploadVideoAction struct {
	Progress int32
}

type TL_sendMessageRecordAudioAction struct {
}

type TL_sendMessageUploadAudioAction struct {
	Progress int32
}

type TL_sendMessageUploadPhotoAction struct {
	Progress int32
}

type TL_sendMessageUploadDocumentAction struct {
	Progress int32
}

type TL_sendMessageGeoLocationAction struct {
}

type TL_sendMessageChooseContactAction struct {
}

type TL_updateServiceNotification struct {
	Flags int32
	// Popup	bool // flags_0?true
	Inbox_date int32
	_Type      string
	Message    string
	Media      TL   // MessageMedia
	Entities   []TL // MessageEntity
}

type TL_userStatusRecently struct {
}

type TL_userStatusLastWeek struct {
}

type TL_userStatusLastMonth struct {
}

type TL_updatePrivacy struct {
	Key   TL   // PrivacyKey
	Rules []TL // PrivacyRule
}

type TL_inputPrivacyKeyStatusTimestamp struct {
}

type TL_privacyKeyStatusTimestamp struct {
}

type TL_inputPrivacyValueAllowContacts struct {
}

type TL_inputPrivacyValueAllowAll struct {
}

type TL_inputPrivacyValueAllowUsers struct {
	Users []TL // InputUser
}

type TL_inputPrivacyValueDisallowContacts struct {
}

type TL_inputPrivacyValueDisallowAll struct {
}

type TL_inputPrivacyValueDisallowUsers struct {
	Users []TL // InputUser
}

type TL_privacyValueAllowContacts struct {
}

type TL_privacyValueAllowAll struct {
}

type TL_privacyValueAllowUsers struct {
	Users []int32
}

type TL_privacyValueDisallowContacts struct {
}

type TL_privacyValueDisallowAll struct {
}

type TL_privacyValueDisallowUsers struct {
	Users []int32
}

type TL_account_privacyRules struct {
	Rules []TL // PrivacyRule
	Users []TL // User
}

type TL_accountDaysTTL struct {
	Days int32
}

type TL_updateUserPhone struct {
	User_id int32
	Phone   string
}

type TL_disabledFeature struct {
	Feature     string
	Description string
}

type TL_documentAttributeImageSize struct {
	W int32
	H int32
}

type TL_documentAttributeAnimated struct {
}

type TL_documentAttributeSticker struct {
	Flags int32
	// Mask	bool // flags_1?true
	Alt         string
	Stickerset  TL // InputStickerSet
	Mask_coords TL // flags_0?MaskCoords
}

type TL_documentAttributeVideo struct {
	Flags int32
	// Round_message	bool // flags_0?true
	Duration int32
	W        int32
	H        int32
}

type TL_documentAttributeAudio struct {
	Flags int32
	// Voice	bool // flags_10?true
	Duration  int32
	Title     string
	Performer string
	Waveform  []byte
}

type TL_documentAttributeFilename struct {
	File_name string
}

type TL_messages_stickersNotModified struct {
}

type TL_messages_stickers struct {
	Hash     string
	Stickers []TL // Document
}

type TL_stickerPack struct {
	Emoticon  string
	Documents []int64
}

type TL_messages_allStickersNotModified struct {
}

type TL_messages_allStickers struct {
	Hash int32
	Sets []TL // StickerSet
}

type TL_account_noPassword struct {
	New_salt                  []byte
	Email_unconfirmed_pattern string
}

type TL_account_password struct {
	Current_salt              []byte
	New_salt                  []byte
	Hint                      string
	Has_recovery              TL // Bool
	Email_unconfirmed_pattern string
}

type TL_updateReadHistoryInbox struct {
	Peer      TL // Peer
	Max_id    int32
	Pts       int32
	Pts_count int32
}

type TL_updateReadHistoryOutbox struct {
	Peer      TL // Peer
	Max_id    int32
	Pts       int32
	Pts_count int32
}

type TL_messages_affectedMessages struct {
	Pts       int32
	Pts_count int32
}

type TL_contactLinkUnknown struct {
}

type TL_contactLinkNone struct {
}

type TL_contactLinkHasPhone struct {
}

type TL_contactLinkContact struct {
}

type TL_updateWebPage struct {
	Webpage   TL // WebPage
	Pts       int32
	Pts_count int32
}

type TL_webPageEmpty struct {
	Id int64
}

type TL_webPagePending struct {
	Id   int64
	Date int32
}

type TL_webPage struct {
	Flags        int32
	Id           int64
	Url          string
	Display_url  string
	Hash         int32
	_Type        string
	Site_name    string
	Title        string
	Description  string
	Photo        TL // flags_4?Photo
	Embed_url    string
	Embed_type   string
	Embed_width  int32
	Embed_height int32
	Duration     int32
	Author       string
	Document     TL // flags_9?Document
	Cached_page  TL // flags_10?Page
}

type TL_messageMediaWebPage struct {
	Webpage TL // WebPage
}

type TL_authorization struct {
	Hash           int64
	Flags          int32
	Device_model   string
	Platform       string
	System_version string
	Api_id         int32
	App_name       string
	App_version    string
	Date_created   int32
	Date_active    int32
	Ip             string
	Country        string
	Region         string
}

type TL_account_authorizations struct {
	Authorizations []TL // Authorization
}

type TL_account_passwordSettings struct {
	Email string
}

type TL_account_passwordInputSettings struct {
	Flags             int32
	New_salt          []byte
	New_password_hash []byte
	Hint              string
	Email             string
}

type TL_auth_passwordRecovery struct {
	Email_pattern string
}

type TL_inputMediaVenue struct {
	Geo_point TL // InputGeoPoint
	Title     string
	Address   string
	Provider  string
	Venue_id  string
}

type TL_messageMediaVenue struct {
	Geo      TL // GeoPoint
	Title    string
	Address  string
	Provider string
	Venue_id string
}

type TL_receivedNotifyMessage struct {
	Id    int32
	Flags int32
}

type TL_chatInviteEmpty struct {
}

type TL_chatInviteExported struct {
	Link string
}

type TL_chatInviteAlready struct {
	Chat TL // Chat
}

type TL_chatInvite struct {
	Flags int32
	// Channel	bool // flags_0?true
	// Broadcast	bool // flags_1?true
	// Public	bool // flags_2?true
	// Megagroup	bool // flags_3?true
	Title              string
	Photo              TL // ChatPhoto
	Participants_count int32
	Participants       []TL // User
}

type TL_messageActionChatJoinedByLink struct {
	Inviter_id int32
}

type TL_updateReadMessagesContents struct {
	Messages  []int32
	Pts       int32
	Pts_count int32
}

type TL_inputStickerSetEmpty struct {
}

type TL_inputStickerSetID struct {
	Id          int64
	Access_hash int64
}

type TL_inputStickerSetShortName struct {
	Short_name string
}

type TL_stickerSet struct {
	Flags int32
	// Installed	bool // flags_0?true
	// Archived	bool // flags_1?true
	// Official	bool // flags_2?true
	// Masks	bool // flags_3?true
	Id          int64
	Access_hash int64
	Title       string
	Short_name  string
	Count       int32
	Hash        int32
}

type TL_messages_stickerSet struct {
	Set       TL   // StickerSet
	Packs     []TL // StickerPack
	Documents []TL // Document
}

type TL_user struct {
	Flags int32
	// Self	bool // flags_10?true
	// Contact	bool // flags_11?true
	// Mutual_contact	bool // flags_12?true
	// Deleted	bool // flags_13?true
	// Bot	bool // flags_14?true
	// Bot_chat_history	bool // flags_15?true
	// Bot_nochats	bool // flags_16?true
	// Verified	bool // flags_17?true
	// Restricted	bool // flags_18?true
	// Min	bool // flags_20?true
	// Bot_inline_geo	bool // flags_21?true
	Id                     int32
	Access_hash            int64
	First_name             string
	Last_name              string
	Username               string
	Phone                  string
	Photo                  TL // flags_5?UserProfilePhoto
	Status                 TL // flags_6?UserStatus
	Bot_info_version       int32
	Restriction_reason     string
	Bot_inline_placeholder string
	Lang_code              string
}

type TL_botCommand struct {
	Command     string
	Description string
}

type TL_botInfo struct {
	User_id     int32
	Description string
	Commands    []TL // BotCommand
}

type TL_keyboardButton struct {
	Text string
}

type TL_keyboardButtonRow struct {
	Buttons []TL // KeyboardButton
}

type TL_replyKeyboardHide struct {
	Flags int32
	// Selective	bool // flags_2?true
}

type TL_replyKeyboardForceReply struct {
	Flags int32
	// Single_use	bool // flags_1?true
	// Selective	bool // flags_2?true
}

type TL_replyKeyboardMarkup struct {
	Flags int32
	// Resize	bool // flags_0?true
	// Single_use	bool // flags_1?true
	// Selective	bool // flags_2?true
	Rows []TL // KeyboardButtonRow
}

type TL_inputMessagesFilterUrl struct {
}

type TL_inputPeerUser struct {
	User_id     int32
	Access_hash int64
}

type TL_inputUser struct {
	User_id     int32
	Access_hash int64
}

type TL_messageEntityUnknown struct {
	Offset int32
	Length int32
}

type TL_messageEntityMention struct {
	Offset int32
	Length int32
}

type TL_messageEntityHashtag struct {
	Offset int32
	Length int32
}

type TL_messageEntityBotCommand struct {
	Offset int32
	Length int32
}

type TL_messageEntityUrl struct {
	Offset int32
	Length int32
}

type TL_messageEntityEmail struct {
	Offset int32
	Length int32
}

type TL_messageEntityBold struct {
	Offset int32
	Length int32
}

type TL_messageEntityItalic struct {
	Offset int32
	Length int32
}

type TL_messageEntityCode struct {
	Offset int32
	Length int32
}

type TL_messageEntityPre struct {
	Offset   int32
	Length   int32
	Language string
}

type TL_messageEntityTextUrl struct {
	Offset int32
	Length int32
	Url    string
}

type TL_updateShortSentMessage struct {
	Flags int32
	// Out	bool // flags_1?true
	Id        int32
	Pts       int32
	Pts_count int32
	Date      int32
	Media     TL   // flags_9?MessageMedia
	Entities  []TL // MessageEntity
}

type TL_inputPeerChannel struct {
	Channel_id  int32
	Access_hash int64
}

type TL_peerChannel struct {
	Channel_id int32
}

type TL_channel struct {
	Flags int32
	// Creator	bool // flags_0?true
	// Left	bool // flags_2?true
	// Editor	bool // flags_3?true
	// Broadcast	bool // flags_5?true
	// Verified	bool // flags_7?true
	// Megagroup	bool // flags_8?true
	// Restricted	bool // flags_9?true
	// Democracy	bool // flags_10?true
	// Signatures	bool // flags_11?true
	// Min	bool // flags_12?true
	Id                 int32
	Access_hash        int64
	Title              string
	Username           string
	Photo              TL // ChatPhoto
	Date               int32
	Version            int32
	Restriction_reason string
	Admin_rights       TL // flags_14?ChannelAdminRights
	Banned_rights      TL // flags_15?ChannelBannedRights
}

type TL_channelForbidden struct {
	Flags int32
	// Broadcast	bool // flags_5?true
	// Megagroup	bool // flags_8?true
	Id          int32
	Access_hash int64
	Title       string
	Until_date  int32
}

type TL_channelFull struct {
	Flags int32
	// Can_view_participants	bool // flags_3?true
	// Can_set_username	bool // flags_6?true
	// Can_set_stickers	bool // flags_7?true
	Id                    int32
	About                 string
	Participants_count    int32
	Admins_count          int32
	Kicked_count          int32
	Banned_count          int32
	Read_inbox_max_id     int32
	Read_outbox_max_id    int32
	Unread_count          int32
	Chat_photo            TL   // Photo
	Notify_settings       TL   // PeerNotifySettings
	Exported_invite       TL   // ExportedChatInvite
	Bot_info              []TL // BotInfo
	Migrated_from_chat_id int32
	Migrated_from_max_id  int32
	Pinned_msg_id         int32
	Stickerset            TL // flags_8?StickerSet
}

type TL_messageActionChannelCreate struct {
	Title string
}

type TL_messages_channelMessages struct {
	Flags    int32
	Pts      int32
	Count    int32
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TL_updateChannelTooLong struct {
	Flags      int32
	Channel_id int32
	Pts        int32
}

type TL_updateChannel struct {
	Channel_id int32
}

type TL_updateNewChannelMessage struct {
	Message   TL // Message
	Pts       int32
	Pts_count int32
}

type TL_updateReadChannelInbox struct {
	Channel_id int32
	Max_id     int32
}

type TL_updateDeleteChannelMessages struct {
	Channel_id int32
	Messages   []int32
	Pts        int32
	Pts_count  int32
}

type TL_updateChannelMessageViews struct {
	Channel_id int32
	Id         int32
	Views      int32
}

type TL_inputChannelEmpty struct {
}

type TL_inputChannel struct {
	Channel_id  int32
	Access_hash int64
}

type TL_contacts_resolvedPeer struct {
	Peer  TL   // Peer
	Chats []TL // Chat
	Users []TL // User
}

type TL_messageRange struct {
	Min_id int32
	Max_id int32
}

type TL_updates_channelDifferenceEmpty struct {
	Flags int32
	// Final	bool // flags_0?true
	Pts     int32
	Timeout int32
}

type TL_updates_channelDifferenceTooLong struct {
	Flags int32
	// Final	bool // flags_0?true
	Pts                   int32
	Timeout               int32
	Top_message           int32
	Read_inbox_max_id     int32
	Read_outbox_max_id    int32
	Unread_count          int32
	Unread_mentions_count int32
	Messages              []TL // Message
	Chats                 []TL // Chat
	Users                 []TL // User
}

type TL_updates_channelDifference struct {
	Flags int32
	// Final	bool // flags_0?true
	Pts           int32
	Timeout       int32
	New_messages  []TL // Message
	Other_updates []TL // Update
	Chats         []TL // Chat
	Users         []TL // User
}

type TL_channelMessagesFilterEmpty struct {
}

type TL_channelMessagesFilter struct {
	Flags int32
	// Exclude_new_messages	bool // flags_1?true
	Ranges []TL // MessageRange
}

type TL_channelParticipant struct {
	User_id int32
	Date    int32
}

type TL_channelParticipantSelf struct {
	User_id    int32
	Inviter_id int32
	Date       int32
}

type TL_channelParticipantCreator struct {
	User_id int32
}

type TL_channelParticipantsRecent struct {
}

type TL_channelParticipantsAdmins struct {
}

type TL_channelParticipantsKicked struct {
	Q string
}

type TL_channels_channelParticipants struct {
	Count        int32
	Participants []TL // ChannelParticipant
	Users        []TL // User
}

type TL_channels_channelParticipant struct {
	Participant TL   // ChannelParticipant
	Users       []TL // User
}

type TL_true struct {
}

type TL_chatParticipantCreator struct {
	User_id int32
}

type TL_chatParticipantAdmin struct {
	User_id    int32
	Inviter_id int32
	Date       int32
}

type TL_updateChatAdmins struct {
	Chat_id int32
	Enabled TL // Bool
	Version int32
}

type TL_updateChatParticipantAdmin struct {
	Chat_id  int32
	User_id  int32
	Is_admin TL // Bool
	Version  int32
}

type TL_messageActionChatMigrateTo struct {
	Channel_id int32
}

type TL_messageActionChannelMigrateFrom struct {
	Title   string
	Chat_id int32
}

type TL_channelParticipantsBots struct {
}

type TL_inputReportReasonSpam struct {
}

type TL_inputReportReasonViolence struct {
}

type TL_inputReportReasonPornography struct {
}

type TL_inputReportReasonOther struct {
	Text string
}

type TL_updateNewStickerSet struct {
	Stickerset TL // messages_StickerSet
}

type TL_updateStickerSetsOrder struct {
	Flags int32
	// Masks	bool // flags_0?true
	Order []int64
}

type TL_updateStickerSets struct {
}

type TL_help_termsOfService struct {
	Text string
}

type TL_foundGif struct {
	Url          string
	Thumb_url    string
	Content_url  string
	Content_type string
	W            int32
	H            int32
}

type TL_inputMediaGifExternal struct {
	Url string
	Q   string
}

type TL_messages_foundGifs struct {
	Next_offset int32
	Results     []TL // FoundGif
}

type TL_inputMessagesFilterGif struct {
}

type TL_updateSavedGifs struct {
}

type TL_updateBotInlineQuery struct {
	Flags    int32
	Query_id int64
	User_id  int32
	Query    string
	Geo      TL // flags_0?GeoPoint
	Offset   string
}

type TL_foundGifCached struct {
	Url      string
	Photo    TL // Photo
	Document TL // Document
}

type TL_messages_savedGifsNotModified struct {
}

type TL_messages_savedGifs struct {
	Hash int32
	Gifs []TL // Document
}

type TL_inputBotInlineMessageMediaAuto struct {
	Flags        int32
	Caption      string
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_inputBotInlineMessageText struct {
	Flags int32
	// No_webpage	bool // flags_0?true
	Message      string
	Entities     []TL // MessageEntity
	Reply_markup TL   // flags_2?ReplyMarkup
}

type TL_inputBotInlineResult struct {
	Flags        int32
	Id           string
	_Type        string
	Title        string
	Description  string
	Url          string
	Thumb_url    string
	Content_url  string
	Content_type string
	W            int32
	H            int32
	Duration     int32
	Send_message TL // InputBotInlineMessage
}

type TL_botInlineMessageMediaAuto struct {
	Flags        int32
	Caption      string
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_botInlineMessageText struct {
	Flags int32
	// No_webpage	bool // flags_0?true
	Message      string
	Entities     []TL // MessageEntity
	Reply_markup TL   // flags_2?ReplyMarkup
}

type TL_botInlineResult struct {
	Flags        int32
	Id           string
	_Type        string
	Title        string
	Description  string
	Url          string
	Thumb_url    string
	Content_url  string
	Content_type string
	W            int32
	H            int32
	Duration     int32
	Send_message TL // BotInlineMessage
}

type TL_messages_botResults struct {
	Flags int32
	// Gallery	bool // flags_0?true
	Query_id    int64
	Next_offset string
	Switch_pm   TL   // flags_2?InlineBotSwitchPM
	Results     []TL // BotInlineResult
	Cache_time  int32
}

type TL_inputMessagesFilterVoice struct {
}

type TL_inputMessagesFilterMusic struct {
}

type TL_updateBotInlineSend struct {
	Flags   int32
	User_id int32
	Query   string
	Geo     TL // flags_0?GeoPoint
	Id      string
	Msg_id  TL // flags_1?InputBotInlineMessageID
}

type TL_inputPrivacyKeyChatInvite struct {
}

type TL_privacyKeyChatInvite struct {
}

type TL_updateEditChannelMessage struct {
	Message   TL // Message
	Pts       int32
	Pts_count int32
}

type TL_exportedMessageLink struct {
	Link string
}

type TL_messageFwdHeader struct {
	Flags        int32
	From_id      int32
	Date         int32
	Channel_id   int32
	Channel_post int32
	Post_author  string
}

type TL_messageActionPinMessage struct {
}

type TL_peerSettings struct {
	Flags int32
	// Report_spam	bool // flags_0?true
}

type TL_updateChannelPinnedMessage struct {
	Channel_id int32
	Id         int32
}

type TL_keyboardButtonUrl struct {
	Text string
	Url  string
}

type TL_keyboardButtonCallback struct {
	Text string
	Data []byte
}

type TL_keyboardButtonRequestPhone struct {
	Text string
}

type TL_keyboardButtonRequestGeoLocation struct {
	Text string
}

type TL_auth_codeTypeSms struct {
}

type TL_auth_codeTypeCall struct {
}

type TL_auth_codeTypeFlashCall struct {
}

type TL_auth_sentCodeTypeApp struct {
	Length int32
}

type TL_auth_sentCodeTypeSms struct {
	Length int32
}

type TL_auth_sentCodeTypeCall struct {
	Length int32
}

type TL_auth_sentCodeTypeFlashCall struct {
	Pattern string
}

type TL_keyboardButtonSwitchInline struct {
	Flags int32
	// Same_peer	bool // flags_0?true
	Text  string
	Query string
}

type TL_replyInlineMarkup struct {
	Rows []TL // KeyboardButtonRow
}

type TL_messages_botCallbackAnswer struct {
	Flags int32
	// Alert	bool // flags_1?true
	// Has_url	bool // flags_3?true
	Message    string
	Url        string
	Cache_time int32
}

type TL_updateBotCallbackQuery struct {
	Flags           int32
	Query_id        int64
	User_id         int32
	Peer            TL // Peer
	Msg_id          int32
	Chat_instance   int64
	Data            []byte
	Game_short_name string
}

type TL_messages_messageEditData struct {
	Flags int32
	// Caption	bool // flags_0?true
}

type TL_updateEditMessage struct {
	Message   TL // Message
	Pts       int32
	Pts_count int32
}

type TL_inputBotInlineMessageMediaGeo struct {
	Flags        int32
	Geo_point    TL // InputGeoPoint
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_inputBotInlineMessageMediaVenue struct {
	Flags        int32
	Geo_point    TL // InputGeoPoint
	Title        string
	Address      string
	Provider     string
	Venue_id     string
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_inputBotInlineMessageMediaContact struct {
	Flags        int32
	Phone_number string
	First_name   string
	Last_name    string
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_botInlineMessageMediaGeo struct {
	Flags        int32
	Geo          TL // GeoPoint
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_botInlineMessageMediaVenue struct {
	Flags        int32
	Geo          TL // GeoPoint
	Title        string
	Address      string
	Provider     string
	Venue_id     string
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_botInlineMessageMediaContact struct {
	Flags        int32
	Phone_number string
	First_name   string
	Last_name    string
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_inputBotInlineResultPhoto struct {
	Id           string
	_Type        string
	Photo        TL // InputPhoto
	Send_message TL // InputBotInlineMessage
}

type TL_inputBotInlineResultDocument struct {
	Flags        int32
	Id           string
	_Type        string
	Title        string
	Description  string
	Document     TL // InputDocument
	Send_message TL // InputBotInlineMessage
}

type TL_botInlineMediaResult struct {
	Flags        int32
	Id           string
	_Type        string
	Photo        TL // flags_0?Photo
	Document     TL // flags_1?Document
	Title        string
	Description  string
	Send_message TL // BotInlineMessage
}

type TL_inputBotInlineMessageID struct {
	Dc_id       int32
	Id          int64
	Access_hash int64
}

type TL_updateInlineBotCallbackQuery struct {
	Flags           int32
	Query_id        int64
	User_id         int32
	Msg_id          TL // InputBotInlineMessageID
	Chat_instance   int64
	Data            []byte
	Game_short_name string
}

type TL_inlineBotSwitchPM struct {
	Text        string
	Start_param string
}

type TL_messageEntityMentionName struct {
	Offset  int32
	Length  int32
	User_id int32
}

type TL_inputMessageEntityMentionName struct {
	Offset  int32
	Length  int32
	User_id TL // InputUser
}

type TL_messages_peerDialogs struct {
	Dialogs  []TL // Dialog
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
	State    TL   // updates_State
}

type TL_topPeer struct {
	Peer   TL // Peer
	Rating float64
}

type TL_topPeerCategoryBotsPM struct {
}

type TL_topPeerCategoryBotsInline struct {
}

type TL_topPeerCategoryCorrespondents struct {
}

type TL_topPeerCategoryGroups struct {
}

type TL_topPeerCategoryChannels struct {
}

type TL_topPeerCategoryPeers struct {
	Category TL // TopPeerCategory
	Count    int32
	Peers    []TL // TopPeer
}

type TL_contacts_topPeersNotModified struct {
}

type TL_contacts_topPeers struct {
	Categories []TL // TopPeerCategoryPeers
	Chats      []TL // Chat
	Users      []TL // User
}

type TL_inputMessagesFilterChatPhotos struct {
}

type TL_updateReadChannelOutbox struct {
	Channel_id int32
	Max_id     int32
}

type TL_updateDraftMessage struct {
	Peer  TL // Peer
	Draft TL // DraftMessage
}

type TL_draftMessageEmpty struct {
}

type TL_draftMessage struct {
	Flags int32
	// No_webpage	bool // flags_1?true
	Reply_to_msg_id int32
	Message         string
	Entities        []TL // MessageEntity
	Date            int32
}

type TL_messageActionHistoryClear struct {
}

type TL_updateReadFeaturedStickers struct {
}

type TL_updateRecentStickers struct {
}

type TL_messages_featuredStickersNotModified struct {
}

type TL_messages_featuredStickers struct {
	Hash   int32
	Sets   []TL // StickerSetCovered
	Unread []int64
}

type TL_messages_recentStickersNotModified struct {
}

type TL_messages_recentStickers struct {
	Hash     int32
	Stickers []TL // Document
}

type TL_messages_archivedStickers struct {
	Count int32
	Sets  []TL // StickerSetCovered
}

type TL_messages_stickerSetInstallResultSuccess struct {
}

type TL_messages_stickerSetInstallResultArchive struct {
	Sets []TL // StickerSetCovered
}

type TL_stickerSetCovered struct {
	Set   TL // StickerSet
	Cover TL // Document
}

type TL_inputMediaPhotoExternal struct {
	Flags       int32
	Url         string
	Caption     string
	Ttl_seconds int32
}

type TL_inputMediaDocumentExternal struct {
	Flags       int32
	Url         string
	Caption     string
	Ttl_seconds int32
}

type TL_updateConfig struct {
}

type TL_updatePtsChanged struct {
}

type TL_messageActionGameScore struct {
	Game_id int64
	Score   int32
}

type TL_documentAttributeHasStickers struct {
}

type TL_keyboardButtonGame struct {
	Text string
}

type TL_stickerSetMultiCovered struct {
	Set    TL   // StickerSet
	Covers []TL // Document
}

type TL_maskCoords struct {
	N    int32
	X    float64
	Y    float64
	Zoom float64
}

type TL_inputStickeredMediaPhoto struct {
	Id TL // InputPhoto
}

type TL_inputStickeredMediaDocument struct {
	Id TL // InputDocument
}

type TL_inputMediaGame struct {
	Id TL // InputGame
}

type TL_messageMediaGame struct {
	Game TL // Game
}

type TL_inputBotInlineMessageGame struct {
	Flags        int32
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_inputBotInlineResultGame struct {
	Id           string
	Short_name   string
	Send_message TL // InputBotInlineMessage
}

type TL_game struct {
	Flags       int32
	Id          int64
	Access_hash int64
	Short_name  string
	Title       string
	Description string
	Photo       TL // Photo
	Document    TL // flags_0?Document
}

type TL_inputGameID struct {
	Id          int64
	Access_hash int64
}

type TL_inputGameShortName struct {
	Bot_id     TL // InputUser
	Short_name string
}

type TL_highScore struct {
	Pos     int32
	User_id int32
	Score   int32
}

type TL_messages_highScores struct {
	Scores []TL // HighScore
	Users  []TL // User
}

type TL_messages_chatsSlice struct {
	Count int32
	Chats []TL // Chat
}

type TL_updateChannelWebPage struct {
	Channel_id int32
	Webpage    TL // WebPage
	Pts        int32
	Pts_count  int32
}

type TL_updates_differenceTooLong struct {
	Pts int32
}

type TL_sendMessageGamePlayAction struct {
}

type TL_webPageNotModified struct {
}

type TL_textEmpty struct {
}

type TL_textPlain struct {
	Text string
}

type TL_textBold struct {
	Text TL // RichText
}

type TL_textItalic struct {
	Text TL // RichText
}

type TL_textUnderline struct {
	Text TL // RichText
}

type TL_textStrike struct {
	Text TL // RichText
}

type TL_textFixed struct {
	Text TL // RichText
}

type TL_textUrl struct {
	Text       TL // RichText
	Url        string
	Webpage_id int64
}

type TL_textEmail struct {
	Text  TL // RichText
	Email string
}

type TL_textConcat struct {
	Texts []TL // RichText
}

type TL_pageBlockTitle struct {
	Text TL // RichText
}

type TL_pageBlockSubtitle struct {
	Text TL // RichText
}

type TL_pageBlockAuthorDate struct {
	Author         TL // RichText
	Published_date int32
}

type TL_pageBlockHeader struct {
	Text TL // RichText
}

type TL_pageBlockSubheader struct {
	Text TL // RichText
}

type TL_pageBlockParagraph struct {
	Text TL // RichText
}

type TL_pageBlockPreformatted struct {
	Text     TL // RichText
	Language string
}

type TL_pageBlockFooter struct {
	Text TL // RichText
}

type TL_pageBlockDivider struct {
}

type TL_pageBlockList struct {
	Ordered TL   // Bool
	Items   []TL // RichText
}

type TL_pageBlockBlockquote struct {
	Text    TL // RichText
	Caption TL // RichText
}

type TL_pageBlockPullquote struct {
	Text    TL // RichText
	Caption TL // RichText
}

type TL_pageBlockPhoto struct {
	Photo_id int64
	Caption  TL // RichText
}

type TL_pageBlockVideo struct {
	Flags int32
	// Autoplay	bool // flags_0?true
	// Loop	bool // flags_1?true
	Video_id int64
	Caption  TL // RichText
}

type TL_pageBlockCover struct {
	Cover TL // PageBlock
}

type TL_pageBlockEmbed struct {
	Flags int32
	// Full_width	bool // flags_0?true
	// Allow_scrolling	bool // flags_3?true
	Url             string
	Html            string
	Poster_photo_id int64
	W               int32
	H               int32
	Caption         TL // RichText
}

type TL_pageBlockEmbedPost struct {
	Url             string
	Webpage_id      int64
	Author_photo_id int64
	Author          string
	Date            int32
	Blocks          []TL // PageBlock
	Caption         TL   // RichText
}

type TL_pageBlockSlideshow struct {
	Items   []TL // PageBlock
	Caption TL   // RichText
}

type TL_pagePart struct {
	Blocks    []TL // PageBlock
	Photos    []TL // Photo
	Documents []TL // Document
}

type TL_pageFull struct {
	Blocks    []TL // PageBlock
	Photos    []TL // Photo
	Documents []TL // Document
}

type TL_updatePhoneCall struct {
	Phone_call TL // PhoneCall
}

type TL_updateDialogPinned struct {
	Flags int32
	// Pinned	bool // flags_0?true
	Peer TL // Peer
}

type TL_updatePinnedDialogs struct {
	Flags int32
	Order []TL // Peer
}

type TL_inputPrivacyKeyPhoneCall struct {
}

type TL_privacyKeyPhoneCall struct {
}

type TL_pageBlockUnsupported struct {
}

type TL_pageBlockAnchor struct {
	Name string
}

type TL_pageBlockCollage struct {
	Items   []TL // PageBlock
	Caption TL   // RichText
}

type TL_inputPhoneCall struct {
	Id          int64
	Access_hash int64
}

type TL_phoneCallEmpty struct {
	Id int64
}

type TL_phoneCallWaiting struct {
	Flags          int32
	Id             int64
	Access_hash    int64
	Date           int32
	Admin_id       int32
	Participant_id int32
	Protocol       TL // PhoneCallProtocol
	Receive_date   int32
}

type TL_phoneCallRequested struct {
	Id             int64
	Access_hash    int64
	Date           int32
	Admin_id       int32
	Participant_id int32
	G_a_hash       []byte
	Protocol       TL // PhoneCallProtocol
}

type TL_phoneCall struct {
	Id                      int64
	Access_hash             int64
	Date                    int32
	Admin_id                int32
	Participant_id          int32
	G_a_or_b                []byte
	Key_fingerprint         int64
	Protocol                TL   // PhoneCallProtocol
	Connection              TL   // PhoneConnection
	Alternative_connections []TL // PhoneConnection
	Start_date              int32
}

type TL_phoneCallDiscarded struct {
	Flags int32
	// Need_rating	bool // flags_2?true
	// Need_debug	bool // flags_3?true
	Id       int64
	Reason   TL // flags_0?PhoneCallDiscardReason
	Duration int32
}

type TL_phoneConnection struct {
	Id       int64
	Ip       string
	Ipv6     string
	Port     int32
	Peer_tag []byte
}

type TL_phoneCallProtocol struct {
	Flags int32
	// Udp_p2p	bool // flags_0?true
	// Udp_reflector	bool // flags_1?true
	Min_layer int32
	Max_layer int32
}

type TL_phone_phoneCall struct {
	Phone_call TL   // PhoneCall
	Users      []TL // User
}

type TL_phoneCallDiscardReasonMissed struct {
}

type TL_phoneCallDiscardReasonDisconnect struct {
}

type TL_phoneCallDiscardReasonHangup struct {
}

type TL_phoneCallDiscardReasonBusy struct {
}

type TL_inputMessagesFilterPhoneCalls struct {
	Flags int32
	// Missed	bool // flags_0?true
}

type TL_messageActionPhoneCall struct {
	Flags    int32
	Call_id  int64
	Reason   TL // flags_0?PhoneCallDiscardReason
	Duration int32
}

type TL_invoice struct {
	Flags int32
	// Test	bool // flags_0?true
	// Name_requested	bool // flags_1?true
	// Phone_requested	bool // flags_2?true
	// Email_requested	bool // flags_3?true
	// Shipping_address_requested	bool // flags_4?true
	// Flexible	bool // flags_5?true
	Currency string
	Prices   []TL // LabeledPrice
}

type TL_inputMediaInvoice struct {
	Flags       int32
	Title       string
	Description string
	Photo       TL // flags_0?InputWebDocument
	Invoice     TL // Invoice
	Payload     []byte
	Provider    string
	Start_param string
}

type TL_messageActionPaymentSentMe struct {
	Flags              int32
	Currency           string
	Total_amount       int64
	Payload            []byte
	Info               TL // flags_0?PaymentRequestedInfo
	Shipping_option_id string
	Charge             TL // PaymentCharge
}

type TL_messageMediaInvoice struct {
	Flags int32
	// Shipping_address_requested	bool // flags_1?true
	// Test	bool // flags_3?true
	Title          string
	Description    string
	Photo          TL // flags_0?WebDocument
	Receipt_msg_id int32
	Currency       string
	Total_amount   int64
	Start_param    string
}

type TL_keyboardButtonBuy struct {
	Text string
}

type TL_messageActionPaymentSent struct {
	Currency     string
	Total_amount int64
}

type TL_payments_paymentForm struct {
	Flags int32
	// Can_save_credentials	bool // flags_2?true
	// Password_missing	bool // flags_3?true
	Bot_id            int32
	Invoice           TL // Invoice
	Provider_id       int32
	Url               string
	Native_provider   string
	Native_params     TL   // flags_4?DataJSON
	Saved_info        TL   // flags_0?PaymentRequestedInfo
	Saved_credentials TL   // flags_1?PaymentSavedCredentials
	Users             []TL // User
}

type TL_postAddress struct {
	Street_line1 string
	Street_line2 string
	City         string
	State        string
	Country_iso2 string
	Post_code    string
}

type TL_paymentRequestedInfo struct {
	Flags            int32
	Name             string
	Phone            string
	Email            string
	Shipping_address TL // flags_3?PostAddress
}

type TL_updateBotWebhookJSON struct {
	Data TL // DataJSON
}

type TL_updateBotWebhookJSONQuery struct {
	Query_id int64
	Data     TL // DataJSON
	Timeout  int32
}

type TL_updateBotShippingQuery struct {
	Query_id         int64
	User_id          int32
	Payload          []byte
	Shipping_address TL // PostAddress
}

type TL_updateBotPrecheckoutQuery struct {
	Flags              int32
	Query_id           int64
	User_id            int32
	Payload            []byte
	Info               TL // flags_0?PaymentRequestedInfo
	Shipping_option_id string
	Currency           string
	Total_amount       int64
}

type TL_dataJSON struct {
	Data string
}

type TL_labeledPrice struct {
	Label  string
	Amount int64
}

type TL_paymentCharge struct {
	Id                 string
	Provider_charge_id string
}

type TL_paymentSavedCredentialsCard struct {
	Id    string
	Title string
}

type TL_webDocument struct {
	Url         string
	Access_hash int64
	Size        int32
	Mime_type   string
	Attributes  []TL // DocumentAttribute
	Dc_id       int32
}

type TL_inputWebDocument struct {
	Url        string
	Size       int32
	Mime_type  string
	Attributes []TL // DocumentAttribute
}

type TL_inputWebFileLocation struct {
	Url         string
	Access_hash int64
}

type TL_upload_webFile struct {
	Size      int32
	Mime_type string
	File_type TL // storage_FileType
	Mtime     int32
	Bytes     []byte
}

type TL_payments_validatedRequestedInfo struct {
	Flags            int32
	Id               string
	Shipping_options []TL // ShippingOption
}

type TL_payments_paymentResult struct {
	Updates TL // Updates
}

type TL_payments_paymentVerficationNeeded struct {
	Url string
}

type TL_payments_paymentReceipt struct {
	Flags             int32
	Date              int32
	Bot_id            int32
	Invoice           TL // Invoice
	Provider_id       int32
	Info              TL // flags_0?PaymentRequestedInfo
	Shipping          TL // flags_1?ShippingOption
	Currency          string
	Total_amount      int64
	Credentials_title string
	Users             []TL // User
}

type TL_payments_savedInfo struct {
	Flags int32
	// Has_saved_credentials	bool // flags_1?true
	Saved_info TL // flags_0?PaymentRequestedInfo
}

type TL_inputPaymentCredentialsSaved struct {
	Id           string
	Tmp_password []byte
}

type TL_inputPaymentCredentials struct {
	Flags int32
	// Save	bool // flags_0?true
	Data TL // DataJSON
}

type TL_account_tmpPassword struct {
	Tmp_password []byte
	Valid_until  int32
}

type TL_shippingOption struct {
	Id     string
	Title  string
	Prices []TL // LabeledPrice
}

type TL_phoneCallAccepted struct {
	Id             int64
	Access_hash    int64
	Date           int32
	Admin_id       int32
	Participant_id int32
	G_b            []byte
	Protocol       TL // PhoneCallProtocol
}

type TL_inputMessagesFilterRoundVoice struct {
}

type TL_inputMessagesFilterRoundVideo struct {
}

type TL_upload_fileCdnRedirect struct {
	Dc_id           int32
	File_token      []byte
	Encryption_key  []byte
	Encryption_iv   []byte
	Cdn_file_hashes []TL // CdnFileHash
}

type TL_sendMessageRecordRoundAction struct {
}

type TL_sendMessageUploadRoundAction struct {
	Progress int32
}

type TL_upload_cdnFileReuploadNeeded struct {
	Request_token []byte
}

type TL_upload_cdnFile struct {
	Bytes []byte
}

type TL_cdnPublicKey struct {
	Dc_id      int32
	Public_key string
}

type TL_cdnConfig struct {
	Public_keys []TL // CdnPublicKey
}

type TL_updateLangPackTooLong struct {
}

type TL_updateLangPack struct {
	Difference TL // LangPackDifference
}

type TL_pageBlockChannel struct {
	Channel TL // Chat
}

type TL_inputStickerSetItem struct {
	Flags       int32
	Document    TL // InputDocument
	Emoji       string
	Mask_coords TL // flags_0?MaskCoords
}

type TL_langPackString struct {
	Key   string
	Value string
}

type TL_langPackStringPluralized struct {
	Flags       int32
	Key         string
	Zero_value  string
	One_value   string
	Two_value   string
	Few_value   string
	Many_value  string
	Other_value string
}

type TL_langPackStringDeleted struct {
	Key string
}

type TL_langPackDifference struct {
	Lang_code    string
	From_version int32
	Version      int32
	Strings      []TL // LangPackString
}

type TL_langPackLanguage struct {
	Name        string
	Native_name string
	Lang_code   string
}

type TL_channelParticipantAdmin struct {
	Flags int32
	// Can_edit	bool // flags_0?true
	User_id      int32
	Inviter_id   int32
	Promoted_by  int32
	Date         int32
	Admin_rights TL // ChannelAdminRights
}

type TL_channelParticipantBanned struct {
	Flags int32
	// Left	bool // flags_0?true
	User_id       int32
	Kicked_by     int32
	Date          int32
	Banned_rights TL // ChannelBannedRights
}

type TL_channelParticipantsBanned struct {
	Q string
}

type TL_channelParticipantsSearch struct {
	Q string
}

type TL_topPeerCategoryPhoneCalls struct {
}

type TL_pageBlockAudio struct {
	Audio_id int64
	Caption  TL // RichText
}

type TL_channelAdminRights struct {
	Flags int32
	// Change_info	bool // flags_0?true
	// Post_messages	bool // flags_1?true
	// Edit_messages	bool // flags_2?true
	// Delete_messages	bool // flags_3?true
	// Ban_users	bool // flags_4?true
	// Invite_users	bool // flags_5?true
	// Invite_link	bool // flags_6?true
	// Pin_messages	bool // flags_7?true
	// Add_admins	bool // flags_9?true
}

type TL_channelBannedRights struct {
	Flags int32
	// View_messages	bool // flags_0?true
	// Send_messages	bool // flags_1?true
	// Send_media	bool // flags_2?true
	// Send_stickers	bool // flags_3?true
	// Send_gifs	bool // flags_4?true
	// Send_games	bool // flags_5?true
	// Send_inline	bool // flags_6?true
	// Embed_links	bool // flags_7?true
	Until_date int32
}

type TL_channelAdminLogEventActionChangeTitle struct {
	Prev_value string
	New_value  string
}

type TL_channelAdminLogEventActionChangeAbout struct {
	Prev_value string
	New_value  string
}

type TL_channelAdminLogEventActionChangeUsername struct {
	Prev_value string
	New_value  string
}

type TL_channelAdminLogEventActionChangePhoto struct {
	Prev_photo TL // ChatPhoto
	New_photo  TL // ChatPhoto
}

type TL_channelAdminLogEventActionToggleInvites struct {
	New_value TL // Bool
}

type TL_channelAdminLogEventActionToggleSignatures struct {
	New_value TL // Bool
}

type TL_channelAdminLogEventActionUpdatePinned struct {
	Message TL // Message
}

type TL_channelAdminLogEventActionEditMessage struct {
	Prev_message TL // Message
	New_message  TL // Message
}

type TL_channelAdminLogEventActionDeleteMessage struct {
	Message TL // Message
}

type TL_channelAdminLogEventActionParticipantJoin struct {
}

type TL_channelAdminLogEventActionParticipantLeave struct {
}

type TL_channelAdminLogEventActionParticipantInvite struct {
	Participant TL // ChannelParticipant
}

type TL_channelAdminLogEventActionParticipantToggleBan struct {
	Prev_participant TL // ChannelParticipant
	New_participant  TL // ChannelParticipant
}

type TL_channelAdminLogEventActionParticipantToggleAdmin struct {
	Prev_participant TL // ChannelParticipant
	New_participant  TL // ChannelParticipant
}

type TL_channelAdminLogEvent struct {
	Id      int64
	Date    int32
	User_id int32
	Action  TL // ChannelAdminLogEventAction
}

type TL_channels_adminLogResults struct {
	Events []TL // ChannelAdminLogEvent
	Chats  []TL // Chat
	Users  []TL // User
}

type TL_channelAdminLogEventsFilter struct {
	Flags int32
	// Join	bool // flags_0?true
	// Leave	bool // flags_1?true
	// Invite	bool // flags_2?true
	// Ban	bool // flags_3?true
	// Unban	bool // flags_4?true
	// Kick	bool // flags_5?true
	// Unkick	bool // flags_6?true
	// Promote	bool // flags_7?true
	// Demote	bool // flags_8?true
	// Info	bool // flags_9?true
	// Settings	bool // flags_10?true
	// Pinned	bool // flags_11?true
	// Edit	bool // flags_12?true
	// Delete	bool // flags_13?true
}

type TL_messageActionScreenshotTaken struct {
}

type TL_popularContact struct {
	Client_id int64
	Importers int32
}

type TL_cdnFileHash struct {
	Offset int32
	Limit  int32
	Hash   []byte
}

type TL_inputMessagesFilterMyMentions struct {
}

type TL_inputMessagesFilterMyMentionsUnread struct {
}

type TL_updateContactsReset struct {
}

type TL_channelAdminLogEventActionChangeStickerSet struct {
	Prev_stickerset TL // InputStickerSet
	New_stickerset  TL // InputStickerSet
}

type TL_updateFavedStickers struct {
}

type TL_messages_favedStickers struct {
	Hash     int32
	Packs    []TL // StickerPack
	Stickers []TL // Document
}

type TL_messages_favedStickersNotModified struct {
}

type TL_updateChannelReadMessagesContents struct {
	Channel_id int32
	Messages   []int32
}

type TL_invokeAfterMsg struct {
	Msg_id int64
	Query  TL
}

type TL_invokeAfterMsgs struct {
	Msg_ids []int64
	Query   TL
}

type TL_auth_checkPhone struct {
	PhoneNumber string
}

type TL_auth_sendCode struct {
	Flags int32
	// Allow_flashcall	bool // flags_0?true
	PhoneNumber   string
	CurrentNumber TL // flags_0?Bool
	ApiID         int32
	ApiHash       string
}

type TL_auth_signUp struct {
	PhoneNumber   string
	PhoneCodeHash string
	PhoneCode     string
	FirstName     string
	LastName      string
}

type TL_auth_signIn struct {
	PhoneNumber   string
	PhoneCodeHash string
	PhoneCode     string
}

type TL_auth_logOut struct {
}

type TL_auth_resetAuthorizations struct {
}

type TL_auth_sendInvites struct {
	Phone_numbers []string
	Message       string
}

type TL_auth_exportAuthorization struct {
	Dc_id int32
}

type TL_auth_importAuthorization struct {
	Id    int32
	Bytes []byte
}

type TL_account_registerDevice struct {
	Token_type int32
	Token      string
}

type TL_account_unregisterDevice struct {
	Token_type int32
	Token      string
}

type TL_account_updateNotifySettings struct {
	Peer     TL // InputNotifyPeer
	Settings TL // InputPeerNotifySettings
}

type TL_account_getNotifySettings struct {
	Peer TL // InputNotifyPeer
}

type TL_account_resetNotifySettings struct {
}

type TL_account_updateProfile struct {
	Flags      int32
	First_name string
	Last_name  string
	About      string
}

type TL_account_updateStatus struct {
	Offline TL // Bool
}

type TL_account_getWallPapers struct {
}

type TL_users_getUsers struct {
	Id []TL // InputUser
}

type TL_users_getFullUser struct {
	Id TL // InputUser
}

type TL_contacts_getStatuses struct {
}

type TL_contacts_getContacts struct {
	Hash int32
}

type TL_contacts_importContacts struct {
	Contacts []TL // InputContact
}

type TL_contacts_search struct {
	Q     string
	Limit int32
}

type TL_contacts_deleteContact struct {
	Id TL // InputUser
}

type TL_contacts_deleteContacts struct {
	Id []TL // InputUser
}

type TL_contacts_block struct {
	Id TL // InputUser
}

type TL_contacts_unblock struct {
	Id TL // InputUser
}

type TL_contacts_getBlocked struct {
	Offset int32
	Limit  int32
}

type TL_messages_getMessages struct {
	Id []int32
}

type TL_messages_getDialogs struct {
	Flags int32
	// Exclude_pinned	bool // flags_0?true
	Offset_date int32
	Offset_id   int32
	Offset_peer TL // InputPeer
	Limit       int32
}

type TL_messages_getHistory struct {
	Peer        TL // InputPeer
	Offset_id   int32
	Offset_date int32
	Add_offset  int32
	Limit       int32
	Max_id      int32
	Min_id      int32
}

type TL_messages_search struct {
	Flags      int32
	Peer       TL // InputPeer
	Q          string
	From_id    TL // flags_0?InputUser
	Filter     TL // MessagesFilter
	Min_date   int32
	Max_date   int32
	Offset_id  int32
	Add_offset int32
	Limit      int32
	Max_id     int32
	Min_id     int32
}

type TL_messages_readHistory struct {
	Peer   TL // InputPeer
	Max_id int32
}

type TL_messages_deleteHistory struct {
	Flags int32
	// Just_clear	bool // flags_0?true
	Peer   TL // InputPeer
	Max_id int32
}

type TL_messages_deleteMessages struct {
	Flags int32
	// Revoke	bool // flags_0?true
	Id []int32
}

type TL_messages_receivedMessages struct {
	Max_id int32
}

type TL_messages_setTyping struct {
	Peer   TL // InputPeer
	Action TL // SendMessageAction
}

type TL_messages_sendMessage struct {
	Flags int32
	// No_webpage	bool // flags_1?true
	// Silent	bool // flags_5?true
	// Background	bool // flags_6?true
	// Clear_draft	bool // flags_7?true
	Peer            TL // InputPeer
	Reply_to_msg_id int32
	Message         string
	Random_id       int64
	Reply_markup    TL   // flags_2?ReplyMarkup
	Entities        []TL // MessageEntity
}

type TL_messages_sendMedia struct {
	Flags int32
	// Silent	bool // flags_5?true
	// Background	bool // flags_6?true
	// Clear_draft	bool // flags_7?true
	Peer            TL // InputPeer
	Reply_to_msg_id int32
	Media           TL // InputMedia
	Random_id       int64
	Reply_markup    TL // flags_2?ReplyMarkup
}

type TL_messages_forwardMessages struct {
	Flags int32
	// Silent	bool // flags_5?true
	// Background	bool // flags_6?true
	// With_my_score	bool // flags_8?true
	From_peer TL // InputPeer
	Id        []int32
	Random_id []int64
	To_peer   TL // InputPeer
}

type TL_messages_getChats struct {
	Id []int32
}

type TL_messages_getFullChat struct {
	Chat_id int32
}

type TL_messages_editChatTitle struct {
	Chat_id int32
	Title   string
}

type TL_messages_editChatPhoto struct {
	Chat_id int32
	Photo   TL // InputChatPhoto
}

type TL_messages_addChatUser struct {
	Chat_id   int32
	User_id   TL // InputUser
	Fwd_limit int32
}

type TL_messages_deleteChatUser struct {
	Chat_id int32
	User_id TL // InputUser
}

type TL_messages_createChat struct {
	Users []TL // InputUser
	Title string
}

type TL_updates_getState struct {
}

type TL_updates_getDifference struct {
	Flags           int32
	Pts             int32
	Pts_total_limit int32
	Date            int32
	Qts             int32
}

type TL_photos_updateProfilePhoto struct {
	Id TL // InputPhoto
}

type TL_photos_uploadProfilePhoto struct {
	File TL // InputFile
}

type TL_upload_saveFilePart struct {
	File_id   int64
	File_part int32
	Bytes     []byte
}

type TL_upload_getFile struct {
	Location TL // InputFileLocation
	Offset   int32
	Limit    int32
}

type TL_help_getConfig struct {
}

type TL_help_getNearestDc struct {
}

type TL_help_getAppUpdate struct {
}

type TL_help_saveAppLog struct {
	Events []TL // InputAppEvent
}

type TL_help_getInviteText struct {
}

type TL_photos_deletePhotos struct {
	Id []TL // InputPhoto
}

type TL_photos_getUserPhotos struct {
	User_id TL // InputUser
	Offset  int32
	Max_id  int64
	Limit   int32
}

type TL_messages_forwardMessage struct {
	Peer      TL // InputPeer
	Id        int32
	Random_id int64
}

type TL_messages_getDhConfig struct {
	Version       int32
	Random_length int32
}

type TL_messages_requestEncryption struct {
	User_id   TL // InputUser
	Random_id int32
	G_a       []byte
}

type TL_messages_acceptEncryption struct {
	Peer            TL // InputEncryptedChat
	G_b             []byte
	Key_fingerprint int64
}

type TL_messages_discardEncryption struct {
	Chat_id int32
}

type TL_messages_setEncryptedTyping struct {
	Peer   TL // InputEncryptedChat
	Typing TL // Bool
}

type TL_messages_readEncryptedHistory struct {
	Peer     TL // InputEncryptedChat
	Max_date int32
}

type TL_messages_sendEncrypted struct {
	Peer      TL // InputEncryptedChat
	Random_id int64
	Data      []byte
}

type TL_messages_sendEncryptedFile struct {
	Peer      TL // InputEncryptedChat
	Random_id int64
	Data      []byte
	File      TL // InputEncryptedFile
}

type TL_messages_sendEncryptedService struct {
	Peer      TL // InputEncryptedChat
	Random_id int64
	Data      []byte
}

type TL_messages_receivedQueue struct {
	Max_qts int32
}

type TL_upload_saveBigFilePart struct {
	File_id          int64
	File_part        int32
	File_total_parts int32
	Bytes            []byte
}

type TL_initConnection struct {
	Api_id           int32
	Device_model     string
	System_version   string
	App_version      string
	System_lang_code string
	Lang_pack        string
	Lang_code        string
	Query            TL
}

type TL_help_getSupport struct {
}

type TL_auth_bindTempAuthKey struct {
	Perm_auth_key_id  int64
	Nonce             int64
	Expires_at        int32
	Encrypted_message []byte
}

type TL_contacts_exportCard struct {
}

type TL_contacts_importCard struct {
	Export_card []int32
}

type TL_messages_readMessageContents struct {
	Id []int32
}

type TL_account_checkUsername struct {
	Username string
}

type TL_account_updateUsername struct {
	Username string
}

type TL_account_getPrivacy struct {
	Key TL // InputPrivacyKey
}

type TL_account_setPrivacy struct {
	Key   TL   // InputPrivacyKey
	Rules []TL // InputPrivacyRule
}

type TL_account_deleteAccount struct {
	Reason string
}

type TL_account_getAccountTTL struct {
}

type TL_account_setAccountTTL struct {
	Ttl TL // AccountDaysTTL
}

type TL_invokeWithLayer struct {
	Layer int32
	Query TL
}

type TL_contacts_resolveUsername struct {
	Username string
}

type TL_account_sendChangePhoneCode struct {
	Flags int32
	// Allow_flashcall	bool // flags_0?true
	Phone_number   string
	Current_number TL // flags_0?Bool
}

type TL_account_changePhone struct {
	Phone_number    string
	Phone_code_hash string
	Phone_code      string
}

type TL_messages_getAllStickers struct {
	Hash int32
}

type TL_account_updateDeviceLocked struct {
	Period int32
}

type TL_account_getPassword struct {
}

type TL_auth_checkPassword struct {
	Password_hash []byte
}

type TL_messages_getWebPagePreview struct {
	Message string
}

type TL_account_getAuthorizations struct {
}

type TL_account_resetAuthorization struct {
	Hash int64
}

type TL_account_getPasswordSettings struct {
	Current_password_hash []byte
}

type TL_account_updatePasswordSettings struct {
	Current_password_hash []byte
	New_settings          TL // account_PasswordInputSettings
}

type TL_auth_requestPasswordRecovery struct {
}

type TL_auth_recoverPassword struct {
	Code string
}

type TL_invokeWithoutUpdates struct {
	Query TL
}

type TL_messages_exportChatInvite struct {
	Chat_id int32
}

type TL_messages_checkChatInvite struct {
	Hash string
}

type TL_messages_importChatInvite struct {
	Hash string
}

type TL_messages_getStickerSet struct {
	Stickerset TL // InputStickerSet
}

type TL_messages_installStickerSet struct {
	Stickerset TL // InputStickerSet
	Archived   TL // Bool
}

type TL_messages_uninstallStickerSet struct {
	Stickerset TL // InputStickerSet
}

type TL_auth_importBotAuthorization struct {
	Flags          int32
	Api_id         int32
	Api_hash       string
	Bot_auth_token string
}

type TL_messages_startBot struct {
	Bot         TL // InputUser
	Peer        TL // InputPeer
	Random_id   int64
	Start_param string
}

type TL_help_getAppChangelog struct {
	Prev_app_version string
}

type TL_messages_reportSpam struct {
	Peer TL // InputPeer
}

type TL_messages_getMessagesViews struct {
	Peer      TL // InputPeer
	Id        []int32
	Increment TL // Bool
}

type TL_updates_getChannelDifference struct {
	Flags int32
	// Force	bool // flags_0?true
	Channel TL // InputChannel
	Filter  TL // ChannelMessagesFilter
	Pts     int32
	Limit   int32
}

type TL_channels_readHistory struct {
	Channel TL // InputChannel
	Max_id  int32
}

type TL_channels_deleteMessages struct {
	Channel TL // InputChannel
	Id      []int32
}

type TL_channels_deleteUserHistory struct {
	Channel TL // InputChannel
	User_id TL // InputUser
}

type TL_channels_reportSpam struct {
	Channel TL // InputChannel
	User_id TL // InputUser
	Id      []int32
}

type TL_channels_getMessages struct {
	Channel TL // InputChannel
	Id      []int32
}

type TL_channels_getParticipants struct {
	Channel TL // InputChannel
	Filter  TL // ChannelParticipantsFilter
	Offset  int32
	Limit   int32
}

type TL_channels_getParticipant struct {
	Channel TL // InputChannel
	User_id TL // InputUser
}

type TL_channels_getChannels struct {
	Id []TL // InputChannel
}

type TL_channels_getFullChannel struct {
	Channel TL // InputChannel
}

type TL_channels_createChannel struct {
	Flags int32
	// Broadcast	bool // flags_0?true
	// Megagroup	bool // flags_1?true
	Title string
	About string
}

type TL_channels_editAbout struct {
	Channel TL // InputChannel
	About   string
}

type TL_channels_editAdmin struct {
	Channel      TL // InputChannel
	User_id      TL // InputUser
	Admin_rights TL // ChannelAdminRights
}

type TL_channels_editTitle struct {
	Channel TL // InputChannel
	Title   string
}

type TL_channels_editPhoto struct {
	Channel TL // InputChannel
	Photo   TL // InputChatPhoto
}

type TL_channels_checkUsername struct {
	Channel  TL // InputChannel
	Username string
}

type TL_channels_updateUsername struct {
	Channel  TL // InputChannel
	Username string
}

type TL_channels_joinChannel struct {
	Channel TL // InputChannel
}

type TL_channels_leaveChannel struct {
	Channel TL // InputChannel
}

type TL_channels_inviteToChannel struct {
	Channel TL   // InputChannel
	Users   []TL // InputUser
}

type TL_channels_exportInvite struct {
	Channel TL // InputChannel
}

type TL_channels_deleteChannel struct {
	Channel TL // InputChannel
}

type TL_messages_toggleChatAdmins struct {
	Chat_id int32
	Enabled TL // Bool
}

type TL_messages_editChatAdmin struct {
	Chat_id  int32
	User_id  TL // InputUser
	Is_admin TL // Bool
}

type TL_messages_migrateChat struct {
	Chat_id int32
}

type TL_messages_searchGlobal struct {
	Q           string
	Offset_date int32
	Offset_peer TL // InputPeer
	Offset_id   int32
	Limit       int32
}

type TL_account_reportPeer struct {
	Peer   TL // InputPeer
	Reason TL // ReportReason
}

type TL_messages_reorderStickerSets struct {
	Flags int32
	// Masks	bool // flags_0?true
	Order []int64
}

type TL_help_getTermsOfService struct {
}

type TL_messages_getDocumentByHash struct {
	Sha256    []byte
	Size      int32
	Mime_type string
}

type TL_messages_searchGifs struct {
	Q      string
	Offset int32
}

type TL_messages_getSavedGifs struct {
	Hash int32
}

type TL_messages_saveGif struct {
	Id     TL // InputDocument
	Unsave TL // Bool
}

type TL_messages_getInlineBotResults struct {
	Flags     int32
	Bot       TL // InputUser
	Peer      TL // InputPeer
	Geo_point TL // flags_0?InputGeoPoint
	Query     string
	Offset    string
}

type TL_messages_setInlineBotResults struct {
	Flags int32
	// Gallery	bool // flags_0?true
	// Private	bool // flags_1?true
	Query_id    int64
	Results     []TL // InputBotInlineResult
	Cache_time  int32
	Next_offset string
	Switch_pm   TL // flags_3?InlineBotSwitchPM
}

type TL_messages_sendInlineBotResult struct {
	Flags int32
	// Silent	bool // flags_5?true
	// Background	bool // flags_6?true
	// Clear_draft	bool // flags_7?true
	Peer            TL // InputPeer
	Reply_to_msg_id int32
	Random_id       int64
	Query_id        int64
	Id              string
}

type TL_channels_toggleInvites struct {
	Channel TL // InputChannel
	Enabled TL // Bool
}

type TL_channels_exportMessageLink struct {
	Channel TL // InputChannel
	Id      int32
}

type TL_channels_toggleSignatures struct {
	Channel TL // InputChannel
	Enabled TL // Bool
}

type TL_messages_hideReportSpam struct {
	Peer TL // InputPeer
}

type TL_messages_getPeerSettings struct {
	Peer TL // InputPeer
}

type TL_channels_updatePinnedMessage struct {
	Flags int32
	// Silent	bool // flags_0?true
	Channel TL // InputChannel
	Id      int32
}

type TL_auth_resendCode struct {
	Phone_number    string
	Phone_code_hash string
}

type TL_auth_cancelCode struct {
	Phone_number    string
	Phone_code_hash string
}

type TL_messages_getMessageEditData struct {
	Peer TL // InputPeer
	Id   int32
}

type TL_messages_editMessage struct {
	Flags int32
	// No_webpage	bool // flags_1?true
	Peer         TL // InputPeer
	Id           int32
	Message      string
	Reply_markup TL   // flags_2?ReplyMarkup
	Entities     []TL // MessageEntity
}

type TL_messages_editInlineBotMessage struct {
	Flags int32
	// No_webpage	bool // flags_1?true
	Id           TL // InputBotInlineMessageID
	Message      string
	Reply_markup TL   // flags_2?ReplyMarkup
	Entities     []TL // MessageEntity
}

type TL_messages_getBotCallbackAnswer struct {
	Flags int32
	// Game	bool // flags_1?true
	Peer   TL // InputPeer
	Msg_id int32
	Data   []byte
}

type TL_messages_setBotCallbackAnswer struct {
	Flags int32
	// Alert	bool // flags_1?true
	Query_id   int64
	Message    string
	Url        string
	Cache_time int32
}

type TL_contacts_getTopPeers struct {
	Flags int32
	// Correspondents	bool // flags_0?true
	// Bots_pm	bool // flags_1?true
	// Bots_inline	bool // flags_2?true
	// Phone_calls	bool // flags_3?true
	// Groups	bool // flags_10?true
	// Channels	bool // flags_15?true
	Offset int32
	Limit  int32
	Hash   int32
}

type TL_contacts_resetTopPeerRating struct {
	Category TL // TopPeerCategory
	Peer     TL // InputPeer
}

type TL_messages_getPeerDialogs struct {
	Peers []TL // InputPeer
}

type TL_messages_saveDraft struct {
	Flags int32
	// No_webpage	bool // flags_1?true
	Reply_to_msg_id int32
	Peer            TL // InputPeer
	Message         string
	Entities        []TL // MessageEntity
}

type TL_messages_getAllDrafts struct {
}

type TL_account_sendConfirmPhoneCode struct {
	Flags int32
	// Allow_flashcall	bool // flags_0?true
	Hash           string
	Current_number TL // flags_0?Bool
}

type TL_account_confirmPhone struct {
	Phone_code_hash string
	Phone_code      string
}

type TL_messages_getFeaturedStickers struct {
	Hash int32
}

type TL_messages_readFeaturedStickers struct {
	Id []int64
}

type TL_messages_getRecentStickers struct {
	Flags int32
	// Attached	bool // flags_0?true
	Hash int32
}

type TL_messages_saveRecentSticker struct {
	Flags int32
	// Attached	bool // flags_0?true
	Id     TL // InputDocument
	Unsave TL // Bool
}

type TL_messages_clearRecentStickers struct {
	Flags int32
	// Attached	bool // flags_0?true
}

type TL_messages_getArchivedStickers struct {
	Flags int32
	// Masks	bool // flags_0?true
	Offset_id int64
	Limit     int32
}

type TL_channels_getAdminedPublicChannels struct {
}

type TL_auth_dropTempAuthKeys struct {
	Except_auth_keys []int64
}

type TL_messages_setGameScore struct {
	Flags int32
	// Edit_message	bool // flags_0?true
	// Force	bool // flags_1?true
	Peer    TL // InputPeer
	Id      int32
	User_id TL // InputUser
	Score   int32
}

type TL_messages_setInlineGameScore struct {
	Flags int32
	// Edit_message	bool // flags_0?true
	// Force	bool // flags_1?true
	Id      TL // InputBotInlineMessageID
	User_id TL // InputUser
	Score   int32
}

type TL_messages_getMaskStickers struct {
	Hash int32
}

type TL_messages_getAttachedStickers struct {
	Media TL // InputStickeredMedia
}

type TL_messages_getGameHighScores struct {
	Peer    TL // InputPeer
	Id      int32
	User_id TL // InputUser
}

type TL_messages_getInlineGameHighScores struct {
	Id      TL // InputBotInlineMessageID
	User_id TL // InputUser
}

type TL_messages_getCommonChats struct {
	User_id TL // InputUser
	Max_id  int32
	Limit   int32
}

type TL_messages_getAllChats struct {
	Except_ids []int32
}

type TL_help_setBotUpdatesStatus struct {
	Pending_updates_count int32
	Message               string
}

type TL_messages_getWebPage struct {
	Url  string
	Hash int32
}

type TL_messages_toggleDialogPin struct {
	Flags int32
	// Pinned	bool // flags_0?true
	Peer TL // InputPeer
}

type TL_messages_reorderPinnedDialogs struct {
	Flags int32
	// Force	bool // flags_0?true
	Order []TL // InputPeer
}

type TL_messages_getPinnedDialogs struct {
}

type TL_phone_requestCall struct {
	User_id   TL // InputUser
	Random_id int32
	G_a_hash  []byte
	Protocol  TL // PhoneCallProtocol
}

type TL_phone_acceptCall struct {
	Peer     TL // InputPhoneCall
	G_b      []byte
	Protocol TL // PhoneCallProtocol
}

type TL_phone_discardCall struct {
	Peer          TL // InputPhoneCall
	Duration      int32
	Reason        TL // PhoneCallDiscardReason
	Connection_id int64
}

type TL_phone_receivedCall struct {
	Peer TL // InputPhoneCall
}

type TL_messages_reportEncryptedSpam struct {
	Peer TL // InputEncryptedChat
}

type TL_payments_getPaymentForm struct {
	Msg_id int32
}

type TL_payments_sendPaymentForm struct {
	Flags              int32
	Msg_id             int32
	Requested_info_id  string
	Shipping_option_id string
	Credentials        TL // InputPaymentCredentials
}

type TL_account_getTmpPassword struct {
	Password_hash []byte
	Period        int32
}

type TL_messages_setBotShippingResults struct {
	Flags            int32
	Query_id         int64
	Error            string
	Shipping_options []TL // ShippingOption
}

type TL_messages_setBotPrecheckoutResults struct {
	Flags int32
	// Success	bool // flags_1?true
	Query_id int64
	Error    string
}

type TL_upload_getWebFile struct {
	Location TL // InputWebFileLocation
	Offset   int32
	Limit    int32
}

type TL_bots_sendCustomRequest struct {
	Custom_method string
	Params        TL // DataJSON
}

type TL_bots_answerWebhookJSONQuery struct {
	Query_id int64
	Data     TL // DataJSON
}

type TL_payments_getPaymentReceipt struct {
	Msg_id int32
}

type TL_payments_validateRequestedInfo struct {
	Flags int32
	// Save	bool // flags_0?true
	Msg_id int32
	Info   TL // PaymentRequestedInfo
}

type TL_payments_getSavedInfo struct {
}

type TL_payments_clearSavedInfo struct {
	Flags int32
	// Credentials	bool // flags_0?true
	// Info	bool // flags_1?true
}

type TL_phone_getCallConfig struct {
}

type TL_phone_confirmCall struct {
	Peer            TL // InputPhoneCall
	G_a             []byte
	Key_fingerprint int64
	Protocol        TL // PhoneCallProtocol
}

type TL_phone_setCallRating struct {
	Peer    TL // InputPhoneCall
	Rating  int32
	Comment string
}

type TL_phone_saveCallDebug struct {
	Peer  TL // InputPhoneCall
	Debug TL // DataJSON
}

type TL_upload_getCdnFile struct {
	File_token []byte
	Offset     int32
	Limit      int32
}

type TL_upload_reuploadCdnFile struct {
	File_token    []byte
	Request_token []byte
}

type TL_help_getCdnConfig struct {
}

type TL_messages_uploadMedia struct {
	Peer  TL // InputPeer
	Media TL // InputMedia
}

type TL_stickers_createStickerSet struct {
	Flags int32
	// Masks	bool // flags_0?true
	User_id    TL // InputUser
	Title      string
	Short_name string
	Stickers   []TL // InputStickerSetItem
}

type TL_langpack_getLangPack struct {
	Lang_code string
}

type TL_langpack_getStrings struct {
	Lang_code string
	Keys      []string
}

type TL_langpack_getDifference struct {
	From_version int32
}

type TL_langpack_getLanguages struct {
}

type TL_channels_editBanned struct {
	Channel       TL // InputChannel
	User_id       TL // InputUser
	Banned_rights TL // ChannelBannedRights
}

type TL_channels_getAdminLog struct {
	Flags         int32
	Channel       TL // InputChannel
	Q             string
	Events_filter TL   // flags_0?ChannelAdminLogEventsFilter
	Admins        []TL // InputUser
	Max_id        int64
	Min_id        int64
	Limit         int32
}

type TL_stickers_removeStickerFromSet struct {
	Sticker TL // InputDocument
}

type TL_stickers_changeStickerPosition struct {
	Sticker  TL // InputDocument
	Position int32
}

type TL_stickers_addStickerToSet struct {
	Stickerset TL // InputStickerSet
	Sticker    TL // InputStickerSetItem
}

type TL_messages_sendScreenshotNotification struct {
	Peer            TL // InputPeer
	Reply_to_msg_id int32
	Random_id       int64
}

type TL_upload_getCdnFileHashes struct {
	File_token []byte
	Offset     int32
}

type TL_messages_getUnreadMentions struct {
	Peer       TL // InputPeer
	Offset_id  int32
	Add_offset int32
	Limit      int32
	Max_id     int32
	Min_id     int32
}

type TL_messages_faveSticker struct {
	Id     TL // InputDocument
	Unfave TL // Bool
}

type TL_channels_setStickers struct {
	Channel    TL // InputChannel
	Stickerset TL // InputStickerSet
}

type TL_contacts_resetSaved struct {
}

type TL_messages_getFavedStickers struct {
	Hash int32
}

type TL_channels_readMessageContents struct {
	Channel TL // InputChannel
	Id      []int32
}

func (e TL_boolFalse) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_boolFalse)
	return x.Buf
}

func (e TL_boolTrue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_boolTrue)
	return x.Buf
}

func (e TL_error) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_error)
	x.Int(e.Code)
	x.String(e.Text)
	return x.Buf
}

func (e TL_null) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_null)
	return x.Buf
}

func (e TL_inputPeerEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPeerEmpty)
	return x.Buf
}

func (e TL_inputPeerSelf) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPeerSelf)
	return x.Buf
}

func (e TL_inputPeerChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPeerChat)
	x.Int(e.Chat_id)
	return x.Buf
}

func (e TL_inputUserEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputUserEmpty)
	return x.Buf
}

func (e TL_inputUserSelf) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputUserSelf)
	return x.Buf
}

func (e TL_inputPhoneContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPhoneContact)
	x.Long(e.Client_id)
	x.String(e.Phone)
	x.String(e.First_name)
	x.String(e.Last_name)
	return x.Buf
}

func (e TL_inputFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputFile)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.String(e.Name)
	x.String(e.Md5_checksum)
	return x.Buf
}

func (e TL_inputMediaEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMediaEmpty)
	return x.Buf
}

func (e TL_inputMediaUploadedPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMediaUploadedPhoto)
	x.Int(e.Flags)
	x.Bytes(e.File.Encode())
	x.String(e.Caption)
	x.Vector(e.Stickers)
	x.Int(e.Ttl_seconds)
	return x.Buf
}

func (e TL_inputMediaPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMediaPhoto)
	x.Int(e.Flags)
	x.Bytes(e.Id.Encode())
	x.String(e.Caption)
	x.Int(e.Ttl_seconds)
	return x.Buf
}

func (e TL_inputMediaGeoPoint) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMediaGeoPoint)
	x.Bytes(e.Geo_point.Encode())
	return x.Buf
}

func (e TL_inputMediaContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMediaContact)
	x.String(e.Phone_number)
	x.String(e.First_name)
	x.String(e.Last_name)
	return x.Buf
}

func (e TL_inputChatPhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputChatPhotoEmpty)
	return x.Buf
}

func (e TL_inputChatUploadedPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputChatUploadedPhoto)
	x.Bytes(e.File.Encode())
	return x.Buf
}

func (e TL_inputChatPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputChatPhoto)
	x.Bytes(e.Id.Encode())
	return x.Buf
}

func (e TL_inputGeoPointEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputGeoPointEmpty)
	return x.Buf
}

func (e TL_inputGeoPoint) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputGeoPoint)
	x.Double(e.Lat)
	x.Double(e.Long)
	return x.Buf
}

func (e TL_inputPhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPhotoEmpty)
	return x.Buf
}

func (e TL_inputPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPhoto)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.Buf
}

func (e TL_inputFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputFileLocation)
	x.Long(e.Volume_id)
	x.Int(e.Local_id)
	x.Long(e.Secret)
	return x.Buf
}

func (e TL_inputAppEvent) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputAppEvent)
	x.Double(e.Time)
	x.String(e._Type)
	x.Long(e.Peer)
	x.String(e.Data)
	return x.Buf
}

func (e TL_peerUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_peerUser)
	x.Int(e.User_id)
	return x.Buf
}

func (e TL_peerChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_peerChat)
	x.Int(e.Chat_id)
	return x.Buf
}

func (e TL_storage_fileUnknown) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_storage_fileUnknown)
	return x.Buf
}

func (e TL_storage_fileJpeg) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_storage_fileJpeg)
	return x.Buf
}

func (e TL_storage_fileGif) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_storage_fileGif)
	return x.Buf
}

func (e TL_storage_filePng) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_storage_filePng)
	return x.Buf
}

func (e TL_storage_fileMp3) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_storage_fileMp3)
	return x.Buf
}

func (e TL_storage_fileMov) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_storage_fileMov)
	return x.Buf
}

func (e TL_storage_filePartial) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_storage_filePartial)
	return x.Buf
}

func (e TL_storage_fileMp4) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_storage_fileMp4)
	return x.Buf
}

func (e TL_storage_fileWebp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_storage_fileWebp)
	return x.Buf
}

func (e TL_fileLocationUnavailable) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_fileLocationUnavailable)
	x.Long(e.Volume_id)
	x.Int(e.Local_id)
	x.Long(e.Secret)
	return x.Buf
}

func (e TL_fileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_fileLocation)
	x.Int(e.Dc_id)
	x.Long(e.Volume_id)
	x.Int(e.Local_id)
	x.Long(e.Secret)
	return x.Buf
}

func (e TL_userEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_userEmpty)
	x.Int(e.Id)
	return x.Buf
}

func (e TL_userProfilePhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_userProfilePhotoEmpty)
	return x.Buf
}

func (e TL_userProfilePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_userProfilePhoto)
	x.Long(e.Photo_id)
	x.Bytes(e.Photo_small.Encode())
	x.Bytes(e.Photo_big.Encode())
	return x.Buf
}

func (e TL_userStatusEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_userStatusEmpty)
	return x.Buf
}

func (e TL_userStatusOnline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_userStatusOnline)
	x.Int(e.Expires)
	return x.Buf
}

func (e TL_userStatusOffline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_userStatusOffline)
	x.Int(e.Was_online)
	return x.Buf
}

func (e TL_chatEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_chatEmpty)
	x.Int(e.Id)
	return x.Buf
}

func (e TL_chat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_chat)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.String(e.Title)
	x.Bytes(e.Photo.Encode())
	x.Int(e.Participants_count)
	x.Int(e.Date)
	x.Int(e.Version)
	x.Bytes(e.Migrated_to.Encode())
	return x.Buf
}

func (e TL_chatForbidden) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_chatForbidden)
	x.Int(e.Id)
	x.String(e.Title)
	return x.Buf
}

func (e TL_chatFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_chatFull)
	x.Int(e.Id)
	x.Bytes(e.Participants.Encode())
	x.Bytes(e.Chat_photo.Encode())
	x.Bytes(e.Notify_settings.Encode())
	x.Bytes(e.Exported_invite.Encode())
	x.Vector(e.Bot_info)
	return x.Buf
}

func (e TL_chatParticipant) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_chatParticipant)
	x.Int(e.User_id)
	x.Int(e.Inviter_id)
	x.Int(e.Date)
	return x.Buf
}

func (e TL_chatParticipantsForbidden) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_chatParticipantsForbidden)
	x.Int(e.Flags)
	x.Int(e.Chat_id)
	x.Bytes(e.Self_participant.Encode())
	return x.Buf
}

func (e TL_chatParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_chatParticipants)
	x.Int(e.Chat_id)
	x.Vector(e.Participants)
	x.Int(e.Version)
	return x.Buf
}

func (e TL_chatPhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_chatPhotoEmpty)
	return x.Buf
}

func (e TL_chatPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_chatPhoto)
	x.Bytes(e.Photo_small.Encode())
	x.Bytes(e.Photo_big.Encode())
	return x.Buf
}

func (e TL_messageEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageEmpty)
	x.Int(e.Id)
	return x.Buf
}

func (e TL_message) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_message)

	var flags uint32 = 0
	if e.Out == true {
		flags |= 1 << 1
	}
	if e.Mentioned == true {
		flags |= 1 << 4
	}
	if e.Media_unread == true {
		flags |= 1 << 5
	}
	if e.Silent == true {
		flags |= 1 << 13
	}
	if e.Post == true {
		flags |= 1 << 14
	}
	if e.From_id != 0 {
		flags |= 1 << 8
	}
	if e.Fwd_from != nil {
		flags |= 1 << 2
	}
	if e.Via_bot_id != 0 {
		flags |= 1 << 11
	}
	if e.Reply_to_msg_id != 0 {
		flags |= 1 << 3
	}
	if e.Media != nil {
		flags |= 1 << 9
	}
	if e.Reply_markup != nil {
		flags |= 1 << 6
	}
	if e.Entities != nil {
		flags |= 1 << 7
	}
	if e.Views != 0 {
		flags |= 1 << 10
	}
	if e.Edit_date != 0 {
		flags |= 1 << 15
	}
	if e.Post_author != "" {
		flags |= 1 << 16
	}
	if e.GroupedId != 0 {
		flags |= 1 << 17
	}
	if e.Likes != 0 {
		flags |= 1 << 18
	}
	if e.Shares != 0 {
		flags |= 1 << 19
	}
	if e.Comments != 0 {
		flags |= 1 << 20
	}

	x.Int(e.Flags)
	x.Int(e.Id)
	if e.From_id != 0 {
		x.Int(e.From_id)
	}
	x.Bytes(e.To_id.Encode())
	if e.Fwd_from != nil {
		x.Bytes(e.Fwd_from.Encode())
	}
	if e.Via_bot_id != 0 {
		x.Int(e.Via_bot_id)
	}
	if e.Reply_to_msg_id != 0 {
		x.Int(e.Reply_to_msg_id)
	}
	x.Int(e.Date)
	x.String(e.Message)
	if e.Media != nil {
		x.Bytes(e.Media.Encode())
	}
	if e.Reply_markup != nil {
		x.Bytes(e.Reply_markup.Encode())
	}
	if e.Entities != nil {
		x.Int(int32(crc32Vector))
		x.Int(int32(len(e.Entities)))
		for _, v := range e.Entities {
			x.Bytes((v).Encode())
		}
	}
	if e.Views != 0 {
		x.Int(e.Views)
	}
	if e.Edit_date != 0 {
		x.Int(e.Edit_date)
	}
	if e.Post_author != "" {
		x.String(e.Post_author)
	}
	if e.GroupedId != 0 {
		x.Long(e.GroupedId)
	}

	if e.Likes != 0 {
		x.Int(e.Likes)
	}
	if e.Shares != 0 {
		x.Int(e.Shares)
	}
	if e.Comments != 0 {
		x.Int(e.Comments)
	}
	return x.Buf
}

func (e TL_messageService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageService)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Int(e.From_id)
	x.Bytes(e.To_id.Encode())
	x.Int(e.Reply_to_msg_id)
	x.Int(e.Date)
	x.Bytes(e.Action.Encode())
	return x.Buf
}

func (e TL_messageMediaEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageMediaEmpty)
	return x.Buf
}

func (e TL_messageMediaPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageMediaPhoto)
	x.Int(e.Flags)
	x.Bytes(e.Photo.Encode())
	x.String(e.Caption)
	x.Int(e.Ttl_seconds)
	return x.Buf
}

func (e TL_messageMediaGeo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageMediaGeo)
	x.Bytes(e.Geo.Encode())
	return x.Buf
}

func (e TL_messageMediaContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageMediaContact)
	x.String(e.Phone_number)
	x.String(e.First_name)
	x.String(e.Last_name)
	x.Int(e.User_id)
	return x.Buf
}

func (e TL_messageMediaUnsupported) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageMediaUnsupported)
	return x.Buf
}

func (e TL_messageActionEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionEmpty)
	return x.Buf
}

func (e TL_messageActionChatCreate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionChatCreate)
	x.String(e.Title)
	x.VectorInt(e.Users)
	return x.Buf
}

func (e TL_messageActionChatEditTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionChatEditTitle)
	x.String(e.Title)
	return x.Buf
}

func (e TL_messageActionChatEditPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionChatEditPhoto)
	x.Bytes(e.Photo.Encode())
	return x.Buf
}

func (e TL_messageActionChatDeletePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionChatDeletePhoto)
	return x.Buf
}

func (e TL_messageActionChatAddUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionChatAddUser)
	x.VectorInt(e.Users)
	return x.Buf
}

func (e TL_messageActionChatDeleteUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionChatDeleteUser)
	x.Int(e.User_id)
	return x.Buf
}

func (e TL_dialog) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_dialog)
	x.Int(e.Flags)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Top_message)
	x.Int(e.Read_inbox_max_id)
	x.Int(e.Read_outbox_max_id)
	x.Int(e.Unread_count)
	x.Int(e.Unread_mentions_count)
	x.Bytes(e.Notify_settings.Encode())
	x.Int(e.Pts)
	x.Bytes(e.Draft.Encode())
	return x.Buf
}

func (e TL_photoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_photoEmpty)
	x.Long(e.Id)
	return x.Buf
}

func (e TL_photo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_photo)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Vector(e.Sizes)
	return x.Buf
}

func (e TL_photoSizeEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_photoSizeEmpty)
	x.String(e._Type)
	return x.Buf
}

func (e TL_photoSize) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_photoSize)
	x.String(e._Type)
	x.Bytes(e.Location.Encode())
	x.Int(e.W)
	x.Int(e.H)
	x.Int(e.Size)
	return x.Buf
}

func (e TL_photoCachedSize) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_photoCachedSize)
	x.String(e._Type)
	x.Bytes(e.Location.Encode())
	x.Int(e.W)
	x.Int(e.H)
	x.StringBytes(e.Bytes)
	return x.Buf
}

func (e TL_geoPointEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_geoPointEmpty)
	return x.Buf
}

func (e TL_geoPoint) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_geoPoint)
	x.Double(e.Long)
	x.Double(e.Lat)
	return x.Buf
}

func (e TL_auth_checkedPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_checkedPhone)
	x.Bytes(e.Phone_registered.Encode())
	return x.Buf
}

func (e TL_auth_sentCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_sentCode)
	x.Int(e.Flags)

	x.Bytes(e.Type.Encode())
	x.String(e.PhoneCodeHash)

	if e.NextType != nil {
		x.Bytes(e.NextType.Encode())
	}
	if e.Timeout != 0 {
		x.Int(e.Timeout)
	}
	if e.TermsOfService != nil {
		x.Bytes(e.TermsOfService.Encode())
	}
	return x.Buf
}

func (e TL_auth_authorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_authorization)
	x.Int(e.Flags)
	x.Int(e.Tmp_sessions)
	x.Bytes(e.User.Encode())
	return x.Buf
}

func (e TL_auth_exportedAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_exportedAuthorization)
	x.Int(e.Id)
	x.StringBytes(e.Bytes)
	return x.Buf
}

func (e TL_inputNotifyPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputNotifyPeer)
	x.Bytes(e.Peer.Encode())
	return x.Buf
}

func (e TL_inputNotifyUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputNotifyUsers)
	return x.Buf
}

func (e TL_inputNotifyChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputNotifyChats)
	return x.Buf
}

func (e TL_inputNotifyAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputNotifyAll)
	return x.Buf
}

func (e TL_inputPeerNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPeerNotifySettings)
	x.Int(e.Flags)
	x.Int(e.Mute_until)
	x.String(e.Sound)
	return x.Buf
}

func (e TL_peerNotifyEventsEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_peerNotifyEventsEmpty)
	return x.Buf
}

func (e TL_peerNotifyEventsAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_peerNotifyEventsAll)
	return x.Buf
}

func (e TL_peerNotifySettingsEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_peerNotifySettingsEmpty)
	return x.Buf
}

func (e TL_peerNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_peerNotifySettings)
	x.Int(e.Flags)
	x.Int(e.Mute_until)
	x.String(e.Sound)
	return x.Buf
}

func (e TL_wallPaper) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_wallPaper)
	x.Int(e.Id)
	x.String(e.Title)
	x.Vector(e.Sizes)
	x.Int(e.Color)
	return x.Buf
}

func (e TL_userFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_userFull)
	x.Int(e.Flags)
	x.Bytes(e.User.Encode())
	x.String(e.About)
	x.Bytes(e.Link.Encode())
	x.Bytes(e.Profile_photo.Encode())
	x.Bytes(e.Notify_settings.Encode())
	x.Bytes(e.Bot_info.Encode())
	x.Int(e.Common_chats_count)
	return x.Buf
}

func (e TL_contact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contact)
	x.Int(e.User_id)
	x.Bytes(e.Mutual.Encode())
	return x.Buf
}

func (e TL_importedContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_importedContact)
	x.Int(e.User_id)
	x.Long(e.Client_id)
	return x.Buf
}

func (e TL_contactBlocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contactBlocked)
	x.Int(e.User_id)
	x.Int(e.Date)
	return x.Buf
}

func (e TL_contactStatus) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contactStatus)
	x.Int(e.User_id)
	x.Bytes(e.Status.Encode())
	return x.Buf
}

func (e TL_contacts_link) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_link)
	x.Bytes(e.My_link.Encode())
	x.Bytes(e.Foreign_link.Encode())
	x.Bytes(e.User.Encode())
	return x.Buf
}

func (e TL_contacts_contacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_contacts)
	x.Vector(e.Contacts)
	x.Int(e.Saved_count)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_vector_wallpaper) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc32Vector)
	x.Vector(e.Wallpapers)
	return x.Buf
}

func (e TL_contacts_contactsNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_contactsNotModified)
	return x.Buf
}

func (e TL_contacts_importedContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_importedContacts)
	x.Vector(e.Imported)
	x.Vector(e.Popular_invites)
	x.VectorLong(e.Retry_contacts)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_contacts_blocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_blocked)
	x.Vector(e.Blocked)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_contacts_blockedSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_blockedSlice)
	x.Int(e.Count)
	x.Vector(e.Blocked)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_contacts_found) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_found)
	x.Vector(e.Results)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_messages_dialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_dialogs)
	x.Vector(e.Dialogs)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_messages_dialogsSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_dialogsSlice)
	x.Int(e.Count)
	x.Vector(e.Dialogs)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_messages_messages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_messages)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_messages_messagesSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_messagesSlice)
	x.Int(e.Count)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_messages_chats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_chats)
	x.Vector(e.Chats)
	return x.Buf
}

func (e TL_messages_chatFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_chatFull)
	x.Bytes(e.Full_chat.Encode())
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_messages_affectedHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_affectedHistory)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	x.Int(e.Offset)
	return x.Buf
}

func (e TL_inputMessagesFilterEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMessagesFilterEmpty)
	return x.Buf
}

func (e TL_inputMessagesFilterPhotos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMessagesFilterPhotos)
	return x.Buf
}

func (e TL_inputMessagesFilterVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMessagesFilterVideo)
	return x.Buf
}

func (e TL_inputMessagesFilterPhotoVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMessagesFilterPhotoVideo)
	return x.Buf
}

func (e TL_updateNewMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateNewMessage)
	x.Bytes(e.Message.Encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.Buf
}

func (e TL_updateMessageID) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateMessageID)
	x.Int(e.Id)
	x.Long(e.Random_id)
	return x.Buf
}

func (e TL_updateDeleteMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateDeleteMessages)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.Buf
}

func (e TL_updateUserTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateUserTyping)
	x.Int(e.User_id)
	x.Bytes(e.Action.Encode())
	return x.Buf
}

func (e TL_updateChatUserTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateChatUserTyping)
	x.Int(e.Chat_id)
	x.Int(e.User_id)
	x.Bytes(e.Action.Encode())
	return x.Buf
}

func (e TL_updateChatParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateChatParticipants)
	x.Bytes(e.Participants.Encode())
	return x.Buf
}

func (e TL_updateUserStatus) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateUserStatus)
	x.Int(e.User_id)
	x.Bytes(e.Status.Encode())
	return x.Buf
}

func (e TL_updateUserName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateUserName)
	x.Int(e.User_id)
	x.String(e.First_name)
	x.String(e.Last_name)
	x.String(e.Username)
	return x.Buf
}

func (e TL_updateUserPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateUserPhoto)
	x.Int(e.User_id)
	x.Int(e.Date)
	x.Bytes(e.Photo.Encode())
	x.Bytes(e.Previous.Encode())
	return x.Buf
}

func (e TL_updateContactRegistered) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateContactRegistered)
	x.Int(e.User_id)
	x.Int(e.Date)
	return x.Buf
}

func (e TL_updateContactLink) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateContactLink)
	x.Int(e.User_id)
	x.Bytes(e.My_link.Encode())
	x.Bytes(e.Foreign_link.Encode())
	return x.Buf
}

func (e TL_updates_state) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updates_state)
	x.Int(e.Pts)
	x.Int(e.Qts)
	x.Int(e.Date)
	x.Int(e.Seq)
	x.Int(e.Unread_count)
	return x.Buf
}

func (e TL_updates_differenceEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updates_differenceEmpty)
	x.Int(e.Date)
	x.Int(e.Seq)
	return x.Buf
}

func (e TL_updates_difference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updates_difference)
	x.Vector(e.New_messages)
	x.Vector(e.New_encrypted_messages)
	x.Vector(e.Other_updates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Bytes(e.State.Encode())
	return x.Buf
}

func (e TL_updates_differenceSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updates_differenceSlice)
	x.Vector(e.New_messages)
	x.Vector(e.New_encrypted_messages)
	x.Vector(e.Other_updates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Bytes(e.Intermediate_state.Encode())
	return x.Buf
}

func (e TL_updatesTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updatesTooLong)
	return x.Buf
}

func (e TL_updateShortMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateShortMessage)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Int(e.User_id)
	x.String(e.Message)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	x.Int(e.Date)
	x.Bytes(e.Fwd_from.Encode())
	x.Int(e.Via_bot_id)
	x.Int(e.Reply_to_msg_id)
	x.Vector(e.Entities)
	return x.Buf
}

func (e TL_updateShortChatMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateShortChatMessage)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Int(e.From_id)
	x.Int(e.Chat_id)
	x.String(e.Message)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	x.Int(e.Date)
	x.Bytes(e.Fwd_from.Encode())
	x.Int(e.Via_bot_id)
	x.Int(e.Reply_to_msg_id)
	x.Vector(e.Entities)
	return x.Buf
}

func (e TL_updateShort) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateShort)
	x.Bytes(e.Update.Encode())
	x.Int(e.Date)
	return x.Buf
}

func (e TL_updatesCombined) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updatesCombined)
	x.Vector(e.Updates)
	x.Vector(e.Users)
	x.Vector(e.Chats)
	x.Int(e.Date)
	x.Int(e.Seq_start)
	x.Int(e.Seq)
	return x.Buf
}

func (e TL_updates) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updates)
	x.Vector(e.Updates)
	x.Vector(e.Users)
	x.Vector(e.Chats)
	x.Int(e.Date)
	x.Int(e.Seq)
	return x.Buf
}

func (e TL_photos_photo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_photos_photo)
	x.Bytes(e.Photo.Encode())
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_upload_file) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_upload_file)
	x.Bytes(e._Type.Encode())
	x.Int(e.MTime)
	x.StringBytes(e.Bytes)
	return x.Buf
}

func (e TL_dcOption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_dcOption)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.String(e.IpAddress)
	x.Int(e.Port)
	if e.Secret != nil {
		x.StringBytes(e.Secret)
	}
	return x.Buf
}

func (e TL_config) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_config)

	// flags
	var flags uint32 = 0
	if e.TmpSessions != 0 {
		flags |= 1 << 0
	}
	if e.SuggestedLangCode != "" {
		flags |= 1 << 2
	}
	if e.LangPackVersion != 0 {
		flags |= 1 << 2
	}
	x.Int(e.Flags)
	x.Int(e.Date)
	x.Int(e.Expires)
	x.Bytes(e.TestMode.Encode())
	x.Int(e.ThisDC)
	x.Vector(e.DcOptions)
	x.Int(e.ChatSizeMax)
	x.Int(e.MegagroupSizeMax)
	x.Int(e.ForwardedCountMax)
	x.Int(e.OnlineUpdatePeriodMs)
	x.Int(e.OfflineBlurTimeoutMs)
	x.Int(e.OfflineIdleTimeoutMs)
	x.Int(e.OnlineCloudTimeoutMs)
	x.Int(e.NotifyCloudDelayMs)
	x.Int(e.NotifyDefaultDelayMs)
	x.Int(e.ChatBigSize)
	x.Int(e.PushChatPeriodMs)
	x.Int(e.PushChatLimit)
	x.Int(e.SavedGifsLimit)
	x.Int(e.EditTimeLimit)
	x.Int(e.RatingEDecay)
	x.Int(e.StickersRecentLimit)
	x.Int(e.StickersFavedLimit)
	x.Int(e.ChannelsReadMediaPeriod)
	if e.TmpSessions != 0 {
		x.Int(e.TmpSessions)
	}
	x.Int(e.PinnedDialogsCountMax)
	x.Int(e.CallReceiveTimeoutMs)
	x.Int(e.CallRingTimeoutMs)
	x.Int(e.CallConnectTimeoutMs)
	x.Int(e.CallPacketTimeoutMs)
	x.String(e.MeUrlPrefix)
	if e.SuggestedLangCode != "" {
		x.String(e.SuggestedLangCode)
	}
	if e.LangPackVersion != 0 {
		x.Int(e.LangPackVersion)
	}
	x.Int(e.Magic)
	x.Int(e.CountDisableFeature)
	return x.Buf
}

func (e TL_nearestDc) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_nearestDc)
	x.String(e.Country)
	x.Int(e.This_dc)
	x.Int(e.Nearest_dc)
	return x.Buf
}

func (e TL_help_appUpdate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_help_appUpdate)
	x.Int(e.Id)
	x.Bytes(e.Critical.Encode())
	x.String(e.Url)
	x.String(e.Text)
	return x.Buf
}

func (e TL_help_noAppUpdate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_help_noAppUpdate)
	return x.Buf
}

func (e TL_help_inviteText) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_help_inviteText)
	x.String(e.Message)
	return x.Buf
}

func (e TL_inputPeerNotifyEventsEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPeerNotifyEventsEmpty)
	return x.Buf
}

func (e TL_inputPeerNotifyEventsAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPeerNotifyEventsAll)
	return x.Buf
}

func (e TL_photos_photos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_photos_photos)
	x.Vector(e.Photos)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_photos_photosSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_photos_photosSlice)
	x.Int(e.Count)
	x.Vector(e.Photos)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_wallPaperSolid) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_wallPaperSolid)
	x.Int(e.Id)
	x.String(e.Title)
	x.Int(e.Bg_color)
	x.Int(e.Color)
	return x.Buf
}

func (e TL_updateNewEncryptedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateNewEncryptedMessage)
	x.Bytes(e.Message.Encode())
	x.Int(e.Qts)
	return x.Buf
}

func (e TL_updateEncryptedChatTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateEncryptedChatTyping)
	x.Int(e.Chat_id)
	return x.Buf
}

func (e TL_updateEncryption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateEncryption)
	x.Bytes(e.Chat.Encode())
	x.Int(e.Date)
	return x.Buf
}

func (e TL_updateEncryptedMessagesRead) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateEncryptedMessagesRead)
	x.Int(e.Chat_id)
	x.Int(e.Max_date)
	x.Int(e.Date)
	return x.Buf
}

func (e TL_encryptedChatEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_encryptedChatEmpty)
	x.Int(e.Id)
	return x.Buf
}

func (e TL_encryptedChatWaiting) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_encryptedChatWaiting)
	x.Int(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	return x.Buf
}

func (e TL_encryptedChatRequested) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_encryptedChatRequested)
	x.Int(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.StringBytes(e.G_a)
	return x.Buf
}

func (e TL_encryptedChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_encryptedChat)
	x.Int(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.StringBytes(e.G_a_or_b)
	x.Long(e.Key_fingerprint)
	return x.Buf
}

func (e TL_encryptedChatDiscarded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_encryptedChatDiscarded)
	x.Int(e.Id)
	return x.Buf
}

func (e TL_inputEncryptedChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputEncryptedChat)
	x.Int(e.Chat_id)
	x.Long(e.Access_hash)
	return x.Buf
}

func (e TL_encryptedFileEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_encryptedFileEmpty)
	return x.Buf
}

func (e TL_encryptedFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_encryptedFile)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Size)
	x.Int(e.Dc_id)
	x.Int(e.Key_fingerprint)
	return x.Buf
}

func (e TL_inputEncryptedFileEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputEncryptedFileEmpty)
	return x.Buf
}

func (e TL_inputEncryptedFileUploaded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputEncryptedFileUploaded)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.String(e.Md5_checksum)
	x.Int(e.Key_fingerprint)
	return x.Buf
}

func (e TL_inputEncryptedFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputEncryptedFile)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.Buf
}

func (e TL_inputEncryptedFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputEncryptedFileLocation)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.Buf
}

func (e TL_encryptedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_encryptedMessage)
	x.Long(e.Random_id)
	x.Int(e.Chat_id)
	x.Int(e.Date)
	x.StringBytes(e.Bytes)
	x.Bytes(e.File.Encode())
	return x.Buf
}

func (e TL_encryptedMessageService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_encryptedMessageService)
	x.Long(e.Random_id)
	x.Int(e.Chat_id)
	x.Int(e.Date)
	x.StringBytes(e.Bytes)
	return x.Buf
}

func (e TL_messages_dhConfigNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_dhConfigNotModified)
	x.StringBytes(e.Random)
	return x.Buf
}

func (e TL_messages_dhConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_dhConfig)
	x.Int(e.G)
	x.StringBytes(e.P)
	x.Int(e.Version)
	x.StringBytes(e.Random)
	return x.Buf
}

func (e TL_messages_sentEncryptedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_sentEncryptedMessage)
	x.Int(e.Date)
	return x.Buf
}

func (e TL_messages_sentEncryptedFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_sentEncryptedFile)
	x.Int(e.Date)
	x.Bytes(e.File.Encode())
	return x.Buf
}

func (e TL_inputFileBig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputFileBig)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.String(e.Name)
	return x.Buf
}

func (e TL_inputEncryptedFileBigUploaded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputEncryptedFileBigUploaded)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.Int(e.Key_fingerprint)
	return x.Buf
}

func (e TL_storage_filePdf) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_storage_filePdf)
	return x.Buf
}

func (e TL_inputMessagesFilterDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMessagesFilterDocument)
	return x.Buf
}

func (e TL_inputMessagesFilterPhotoVideoDocuments) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMessagesFilterPhotoVideoDocuments)
	return x.Buf
}

func (e TL_updateChatParticipantAdd) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateChatParticipantAdd)
	x.Int(e.Chat_id)
	x.Int(e.User_id)
	x.Int(e.Inviter_id)
	x.Int(e.Date)
	x.Int(e.Version)
	return x.Buf
}

func (e TL_updateChatParticipantDelete) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateChatParticipantDelete)
	x.Int(e.Chat_id)
	x.Int(e.User_id)
	x.Int(e.Version)
	return x.Buf
}

func (e TL_updateDcOptions) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateDcOptions)
	x.Vector(e.Dc_options)
	return x.Buf
}

func (e TL_inputMediaUploadedDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMediaUploadedDocument)
	x.Int(e.Flags)
	x.Bytes(e.File.Encode())
	x.Bytes(e.Thumb.Encode())
	x.String(e.Mime_type)
	x.Vector(e.Attributes)
	x.String(e.Caption)
	x.Vector(e.Stickers)
	x.Int(e.Ttl_seconds)
	return x.Buf
}

func (e TL_inputMediaDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMediaDocument)
	x.Int(e.Flags)
	x.Bytes(e.Id.Encode())
	x.String(e.Caption)
	x.Int(e.Ttl_seconds)
	return x.Buf
}

func (e TL_messageMediaDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageMediaDocument)
	x.Int(e.Flags)
	x.Bytes(e.Document.Encode())
	x.String(e.Caption)
	x.Int(e.Ttl_seconds)
	return x.Buf
}

func (e TL_inputDocumentEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputDocumentEmpty)
	return x.Buf
}

func (e TL_inputDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputDocument)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.Buf
}

func (e TL_inputDocumentFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputDocumentFileLocation)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Version)
	return x.Buf
}

func (e TL_documentEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_documentEmpty)
	x.Long(e.Id)
	return x.Buf
}

func (e TL_document) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_document)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.String(e.Mime_type)
	x.Int(e.Size)
	x.Bytes(e.Thumb.Encode())
	x.Int(e.Dc_id)
	x.Int(e.Version)
	x.Vector(e.Attributes)
	return x.Buf
}

func (e TL_help_support) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_help_support)
	x.String(e.Phone_number)
	x.Bytes(e.User.Encode())
	return x.Buf
}

func (e TL_notifyAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_notifyAll)
	return x.Buf
}

func (e TL_notifyChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_notifyChats)
	return x.Buf
}

func (e TL_notifyPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_notifyPeer)
	x.Bytes(e.Peer.Encode())
	return x.Buf
}

func (e TL_notifyUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_notifyUsers)
	return x.Buf
}

func (e TL_updateUserBlocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateUserBlocked)
	x.Int(e.User_id)
	x.Bytes(e.Blocked.Encode())
	return x.Buf
}

func (e TL_updateNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateNotifySettings)
	x.Bytes(e.Peer.Encode())
	x.Bytes(e.Notify_settings.Encode())
	return x.Buf
}

func (e TL_sendMessageTypingAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_sendMessageTypingAction)
	return x.Buf
}

func (e TL_sendMessageCancelAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_sendMessageCancelAction)
	return x.Buf
}

func (e TL_sendMessageRecordVideoAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_sendMessageRecordVideoAction)
	return x.Buf
}

func (e TL_sendMessageUploadVideoAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_sendMessageUploadVideoAction)
	x.Int(e.Progress)
	return x.Buf
}

func (e TL_sendMessageRecordAudioAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_sendMessageRecordAudioAction)
	return x.Buf
}

func (e TL_sendMessageUploadAudioAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_sendMessageUploadAudioAction)
	x.Int(e.Progress)
	return x.Buf
}

func (e TL_sendMessageUploadPhotoAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_sendMessageUploadPhotoAction)
	x.Int(e.Progress)
	return x.Buf
}

func (e TL_sendMessageUploadDocumentAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_sendMessageUploadDocumentAction)
	x.Int(e.Progress)
	return x.Buf
}

func (e TL_sendMessageGeoLocationAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_sendMessageGeoLocationAction)
	return x.Buf
}

func (e TL_sendMessageChooseContactAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_sendMessageChooseContactAction)
	return x.Buf
}

func (e TL_updateServiceNotification) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateServiceNotification)
	x.Int(e.Flags)
	x.Int(e.Inbox_date)
	x.String(e._Type)
	x.String(e.Message)
	x.Bytes(e.Media.Encode())
	x.Vector(e.Entities)
	return x.Buf
}

func (e TL_userStatusRecently) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_userStatusRecently)
	return x.Buf
}

func (e TL_userStatusLastWeek) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_userStatusLastWeek)
	return x.Buf
}

func (e TL_userStatusLastMonth) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_userStatusLastMonth)
	return x.Buf
}

func (e TL_updatePrivacy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updatePrivacy)
	x.Bytes(e.Key.Encode())
	x.Vector(e.Rules)
	return x.Buf
}

func (e TL_inputPrivacyKeyStatusTimestamp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPrivacyKeyStatusTimestamp)
	return x.Buf
}

func (e TL_privacyKeyStatusTimestamp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_privacyKeyStatusTimestamp)
	return x.Buf
}

func (e TL_inputPrivacyValueAllowContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPrivacyValueAllowContacts)
	return x.Buf
}

func (e TL_inputPrivacyValueAllowAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPrivacyValueAllowAll)
	return x.Buf
}

func (e TL_inputPrivacyValueAllowUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPrivacyValueAllowUsers)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_inputPrivacyValueDisallowContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPrivacyValueDisallowContacts)
	return x.Buf
}

func (e TL_inputPrivacyValueDisallowAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPrivacyValueDisallowAll)
	return x.Buf
}

func (e TL_inputPrivacyValueDisallowUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPrivacyValueDisallowUsers)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_privacyValueAllowContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_privacyValueAllowContacts)
	return x.Buf
}

func (e TL_privacyValueAllowAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_privacyValueAllowAll)
	return x.Buf
}

func (e TL_privacyValueAllowUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_privacyValueAllowUsers)
	x.VectorInt(e.Users)
	return x.Buf
}

func (e TL_privacyValueDisallowContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_privacyValueDisallowContacts)
	return x.Buf
}

func (e TL_privacyValueDisallowAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_privacyValueDisallowAll)
	return x.Buf
}

func (e TL_privacyValueDisallowUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_privacyValueDisallowUsers)
	x.VectorInt(e.Users)
	return x.Buf
}

func (e TL_account_privacyRules) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_privacyRules)
	x.Vector(e.Rules)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_accountDaysTTL) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_accountDaysTTL)
	x.Int(e.Days)
	return x.Buf
}

func (e TL_updateUserPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateUserPhone)
	x.Int(e.User_id)
	x.String(e.Phone)
	return x.Buf
}

func (e TL_disabledFeature) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_disabledFeature)
	x.String(e.Feature)
	x.String(e.Description)
	return x.Buf
}

func (e TL_documentAttributeImageSize) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_documentAttributeImageSize)
	x.Int(e.W)
	x.Int(e.H)
	return x.Buf
}

func (e TL_documentAttributeAnimated) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_documentAttributeAnimated)
	return x.Buf
}

func (e TL_documentAttributeSticker) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_documentAttributeSticker)
	x.Int(e.Flags)
	x.String(e.Alt)
	x.Bytes(e.Stickerset.Encode())
	x.Bytes(e.Mask_coords.Encode())
	return x.Buf
}

func (e TL_documentAttributeVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_documentAttributeVideo)
	x.Int(e.Flags)
	x.Int(e.Duration)
	x.Int(e.W)
	x.Int(e.H)
	return x.Buf
}

func (e TL_documentAttributeAudio) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_documentAttributeAudio)
	x.Int(e.Flags)
	x.Int(e.Duration)
	x.String(e.Title)
	x.String(e.Performer)
	x.StringBytes(e.Waveform)
	return x.Buf
}

func (e TL_documentAttributeFilename) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_documentAttributeFilename)
	x.String(e.File_name)
	return x.Buf
}

func (e TL_messages_stickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_stickersNotModified)
	return x.Buf
}

func (e TL_messages_stickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_stickers)
	x.String(e.Hash)
	x.Vector(e.Stickers)
	return x.Buf
}

func (e TL_stickerPack) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_stickerPack)
	x.String(e.Emoticon)
	x.VectorLong(e.Documents)
	return x.Buf
}

func (e TL_messages_allStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_allStickersNotModified)
	return x.Buf
}

func (e TL_messages_allStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_allStickers)
	x.Int(e.Hash)
	x.Vector(e.Sets)
	return x.Buf
}

func (e TL_account_noPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_noPassword)
	x.StringBytes(e.New_salt)
	x.String(e.Email_unconfirmed_pattern)
	return x.Buf
}

func (e TL_account_password) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_password)
	x.StringBytes(e.Current_salt)
	x.StringBytes(e.New_salt)
	x.String(e.Hint)
	x.Bytes(e.Has_recovery.Encode())
	x.String(e.Email_unconfirmed_pattern)
	return x.Buf
}

func (e TL_updateReadHistoryInbox) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateReadHistoryInbox)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Max_id)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.Buf
}

func (e TL_updateReadHistoryOutbox) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateReadHistoryOutbox)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Max_id)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.Buf
}

func (e TL_messages_affectedMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_affectedMessages)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.Buf
}

func (e TL_contactLinkUnknown) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contactLinkUnknown)
	return x.Buf
}

func (e TL_contactLinkNone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contactLinkNone)
	return x.Buf
}

func (e TL_contactLinkHasPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contactLinkHasPhone)
	return x.Buf
}

func (e TL_contactLinkContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contactLinkContact)
	return x.Buf
}

func (e TL_updateWebPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateWebPage)
	x.Bytes(e.Webpage.Encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.Buf
}

func (e TL_webPageEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_webPageEmpty)
	x.Long(e.Id)
	return x.Buf
}

func (e TL_webPagePending) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_webPagePending)
	x.Long(e.Id)
	x.Int(e.Date)
	return x.Buf
}

func (e TL_webPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_webPage)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.String(e.Url)
	x.String(e.Display_url)
	x.Int(e.Hash)
	x.String(e._Type)
	x.String(e.Site_name)
	x.String(e.Title)
	x.String(e.Description)
	x.Bytes(e.Photo.Encode())
	x.String(e.Embed_url)
	x.String(e.Embed_type)
	x.Int(e.Embed_width)
	x.Int(e.Embed_height)
	x.Int(e.Duration)
	x.String(e.Author)
	x.Bytes(e.Document.Encode())
	x.Bytes(e.Cached_page.Encode())
	return x.Buf
}

func (e TL_messageMediaWebPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageMediaWebPage)
	x.Bytes(e.Webpage.Encode())
	return x.Buf
}

func (e TL_authorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_authorization)
	x.Long(e.Hash)
	x.Int(e.Flags)
	x.String(e.Device_model)
	x.String(e.Platform)
	x.String(e.System_version)
	x.Int(e.Api_id)
	x.String(e.App_name)
	x.String(e.App_version)
	x.Int(e.Date_created)
	x.Int(e.Date_active)
	x.String(e.Ip)
	x.String(e.Country)
	x.String(e.Region)
	return x.Buf
}

func (e TL_account_authorizations) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_authorizations)
	x.Vector(e.Authorizations)
	return x.Buf
}

func (e TL_account_passwordSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_passwordSettings)
	x.String(e.Email)
	return x.Buf
}

func (e TL_account_passwordInputSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_passwordInputSettings)
	x.Int(e.Flags)
	x.StringBytes(e.New_salt)
	x.StringBytes(e.New_password_hash)
	x.String(e.Hint)
	x.String(e.Email)
	return x.Buf
}

func (e TL_auth_passwordRecovery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_passwordRecovery)
	x.String(e.Email_pattern)
	return x.Buf
}

func (e TL_inputMediaVenue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMediaVenue)
	x.Bytes(e.Geo_point.Encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.Venue_id)
	return x.Buf
}

func (e TL_messageMediaVenue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageMediaVenue)
	x.Bytes(e.Geo.Encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.Venue_id)
	return x.Buf
}

func (e TL_receivedNotifyMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_receivedNotifyMessage)
	x.Int(e.Id)
	x.Int(e.Flags)
	return x.Buf
}

func (e TL_chatInviteEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_chatInviteEmpty)
	return x.Buf
}

func (e TL_chatInviteExported) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_chatInviteExported)
	x.String(e.Link)
	return x.Buf
}

func (e TL_chatInviteAlready) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_chatInviteAlready)
	x.Bytes(e.Chat.Encode())
	return x.Buf
}

func (e TL_chatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_chatInvite)
	x.Int(e.Flags)
	x.String(e.Title)
	x.Bytes(e.Photo.Encode())
	x.Int(e.Participants_count)
	x.Vector(e.Participants)
	return x.Buf
}

func (e TL_messageActionChatJoinedByLink) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionChatJoinedByLink)
	x.Int(e.Inviter_id)
	return x.Buf
}

func (e TL_updateReadMessagesContents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateReadMessagesContents)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.Buf
}

func (e TL_inputStickerSetEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputStickerSetEmpty)
	return x.Buf
}

func (e TL_inputStickerSetID) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputStickerSetID)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.Buf
}

func (e TL_inputStickerSetShortName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputStickerSetShortName)
	x.String(e.Short_name)
	return x.Buf
}

func (e TL_stickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_stickerSet)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.String(e.Title)
	x.String(e.Short_name)
	x.Int(e.Count)
	x.Int(e.Hash)
	return x.Buf
}

func (e TL_messages_stickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_stickerSet)
	x.Bytes(e.Set.Encode())
	x.Vector(e.Packs)
	x.Vector(e.Documents)
	return x.Buf
}

func (e TL_user) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_user)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Long(e.Access_hash)
	x.String(e.First_name)
	x.String(e.Last_name)
	x.String(e.Username)
	x.String(e.Phone)
	x.Bytes(e.Photo.Encode())
	x.Bytes(e.Status.Encode())
	x.Int(e.Bot_info_version)
	x.String(e.Restriction_reason)
	x.String(e.Bot_inline_placeholder)
	x.String(e.Lang_code)
	return x.Buf
}

func (e TL_botCommand) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_botCommand)
	x.String(e.Command)
	x.String(e.Description)
	return x.Buf
}

func (e TL_botInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_botInfo)
	x.Int(e.User_id)
	x.String(e.Description)
	x.Vector(e.Commands)
	return x.Buf
}

func (e TL_keyboardButton) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_keyboardButton)
	x.String(e.Text)
	return x.Buf
}

func (e TL_keyboardButtonRow) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_keyboardButtonRow)
	x.Vector(e.Buttons)
	return x.Buf
}

func (e TL_replyKeyboardHide) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_replyKeyboardHide)
	x.Int(e.Flags)
	return x.Buf
}

func (e TL_replyKeyboardForceReply) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_replyKeyboardForceReply)
	x.Int(e.Flags)
	return x.Buf
}

func (e TL_replyKeyboardMarkup) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_replyKeyboardMarkup)
	x.Int(e.Flags)
	x.Vector(e.Rows)
	return x.Buf
}

func (e TL_inputMessagesFilterUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMessagesFilterUrl)
	return x.Buf
}

func (e TL_inputPeerUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPeerUser)
	x.Int(e.User_id)
	x.Long(e.Access_hash)
	return x.Buf
}

func (e TL_inputUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputUser)
	x.Int(e.User_id)
	x.Long(e.Access_hash)
	return x.Buf
}

func (e TL_messageEntityUnknown) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageEntityUnknown)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.Buf
}

func (e TL_messageEntityMention) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageEntityMention)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.Buf
}

func (e TL_messageEntityHashtag) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageEntityHashtag)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.Buf
}

func (e TL_messageEntityBotCommand) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageEntityBotCommand)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.Buf
}

func (e TL_messageEntityUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageEntityUrl)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.Buf
}

func (e TL_messageEntityEmail) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageEntityEmail)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.Buf
}

func (e TL_messageEntityBold) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageEntityBold)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.Buf
}

func (e TL_messageEntityItalic) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageEntityItalic)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.Buf
}

func (e TL_messageEntityCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageEntityCode)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.Buf
}

func (e TL_messageEntityPre) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageEntityPre)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.String(e.Language)
	return x.Buf
}

func (e TL_messageEntityTextUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageEntityTextUrl)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.String(e.Url)
	return x.Buf
}

func (e TL_updateShortSentMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateShortSentMessage)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	x.Int(e.Date)
	x.Bytes(e.Media.Encode())
	x.Vector(e.Entities)
	return x.Buf
}

func (e TL_inputPeerChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPeerChannel)
	x.Int(e.Channel_id)
	x.Long(e.Access_hash)
	return x.Buf
}

func (e TL_peerChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_peerChannel)
	x.Int(e.Channel_id)
	return x.Buf
}

func (e TL_channel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channel)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Long(e.Access_hash)
	x.String(e.Title)
	x.String(e.Username)
	x.Bytes(e.Photo.Encode())
	x.Int(e.Date)
	x.Int(e.Version)
	x.String(e.Restriction_reason)
	x.Bytes(e.Admin_rights.Encode())
	x.Bytes(e.Banned_rights.Encode())
	return x.Buf
}

func (e TL_channelForbidden) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelForbidden)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Long(e.Access_hash)
	x.String(e.Title)
	x.Int(e.Until_date)
	return x.Buf
}

func (e TL_channelFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelFull)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.String(e.About)
	x.Int(e.Participants_count)
	x.Int(e.Admins_count)
	x.Int(e.Kicked_count)
	x.Int(e.Banned_count)
	x.Int(e.Read_inbox_max_id)
	x.Int(e.Read_outbox_max_id)
	x.Int(e.Unread_count)
	x.Bytes(e.Chat_photo.Encode())
	x.Bytes(e.Notify_settings.Encode())
	x.Bytes(e.Exported_invite.Encode())
	x.Vector(e.Bot_info)
	x.Int(e.Migrated_from_chat_id)
	x.Int(e.Migrated_from_max_id)
	x.Int(e.Pinned_msg_id)
	x.Bytes(e.Stickerset.Encode())
	return x.Buf
}

func (e TL_messageActionChannelCreate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionChannelCreate)
	x.String(e.Title)
	return x.Buf
}

func (e TL_messages_channelMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_channelMessages)
	x.Int(e.Flags)
	x.Int(e.Pts)
	x.Int(e.Count)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_updateChannelTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateChannelTooLong)
	x.Int(e.Flags)
	x.Int(e.Channel_id)
	x.Int(e.Pts)
	return x.Buf
}

func (e TL_updateChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateChannel)
	x.Int(e.Channel_id)
	return x.Buf
}

func (e TL_updateNewChannelMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateNewChannelMessage)
	x.Bytes(e.Message.Encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.Buf
}

func (e TL_updateReadChannelInbox) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateReadChannelInbox)
	x.Int(e.Channel_id)
	x.Int(e.Max_id)
	return x.Buf
}

func (e TL_updateDeleteChannelMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateDeleteChannelMessages)
	x.Int(e.Channel_id)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.Buf
}

func (e TL_updateChannelMessageViews) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateChannelMessageViews)
	x.Int(e.Channel_id)
	x.Int(e.Id)
	x.Int(e.Views)
	return x.Buf
}

func (e TL_inputChannelEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputChannelEmpty)
	return x.Buf
}

func (e TL_inputChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputChannel)
	x.Int(e.Channel_id)
	x.Long(e.Access_hash)
	return x.Buf
}

func (e TL_contacts_resolvedPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_resolvedPeer)
	x.Bytes(e.Peer.Encode())
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_messageRange) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageRange)
	x.Int(e.Min_id)
	x.Int(e.Max_id)
	return x.Buf
}

func (e TL_updates_channelDifferenceEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updates_channelDifferenceEmpty)
	x.Int(e.Flags)
	x.Int(e.Pts)
	x.Int(e.Timeout)
	return x.Buf
}

func (e TL_updates_channelDifferenceTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updates_channelDifferenceTooLong)
	x.Int(e.Flags)
	x.Int(e.Pts)
	x.Int(e.Timeout)
	x.Int(e.Top_message)
	x.Int(e.Read_inbox_max_id)
	x.Int(e.Read_outbox_max_id)
	x.Int(e.Unread_count)
	x.Int(e.Unread_mentions_count)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_updates_channelDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updates_channelDifference)
	x.Int(e.Flags)
	x.Int(e.Pts)
	x.Int(e.Timeout)
	x.Vector(e.New_messages)
	x.Vector(e.Other_updates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_channelMessagesFilterEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelMessagesFilterEmpty)
	return x.Buf
}

func (e TL_channelMessagesFilter) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelMessagesFilter)
	x.Int(e.Flags)
	x.Vector(e.Ranges)
	return x.Buf
}

func (e TL_channelParticipant) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelParticipant)
	x.Int(e.User_id)
	x.Int(e.Date)
	return x.Buf
}

func (e TL_channelParticipantSelf) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelParticipantSelf)
	x.Int(e.User_id)
	x.Int(e.Inviter_id)
	x.Int(e.Date)
	return x.Buf
}

func (e TL_channelParticipantCreator) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelParticipantCreator)
	x.Int(e.User_id)
	return x.Buf
}

func (e TL_channelParticipantsRecent) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelParticipantsRecent)
	return x.Buf
}

func (e TL_channelParticipantsAdmins) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelParticipantsAdmins)
	return x.Buf
}

func (e TL_channelParticipantsKicked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelParticipantsKicked)
	x.String(e.Q)
	return x.Buf
}

func (e TL_channels_channelParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_channelParticipants)
	x.Int(e.Count)
	x.Vector(e.Participants)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_channels_channelParticipant) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_channelParticipant)
	x.Bytes(e.Participant.Encode())
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_true) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_true)
	return x.Buf
}

func (e TL_chatParticipantCreator) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_chatParticipantCreator)
	x.Int(e.User_id)
	return x.Buf
}

func (e TL_chatParticipantAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_chatParticipantAdmin)
	x.Int(e.User_id)
	x.Int(e.Inviter_id)
	x.Int(e.Date)
	return x.Buf
}

func (e TL_updateChatAdmins) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateChatAdmins)
	x.Int(e.Chat_id)
	x.Bytes(e.Enabled.Encode())
	x.Int(e.Version)
	return x.Buf
}

func (e TL_updateChatParticipantAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateChatParticipantAdmin)
	x.Int(e.Chat_id)
	x.Int(e.User_id)
	x.Bytes(e.Is_admin.Encode())
	x.Int(e.Version)
	return x.Buf
}

func (e TL_messageActionChatMigrateTo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionChatMigrateTo)
	x.Int(e.Channel_id)
	return x.Buf
}

func (e TL_messageActionChannelMigrateFrom) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionChannelMigrateFrom)
	x.String(e.Title)
	x.Int(e.Chat_id)
	return x.Buf
}

func (e TL_channelParticipantsBots) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelParticipantsBots)
	return x.Buf
}

func (e TL_inputReportReasonSpam) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputReportReasonSpam)
	return x.Buf
}

func (e TL_inputReportReasonViolence) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputReportReasonViolence)
	return x.Buf
}

func (e TL_inputReportReasonPornography) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputReportReasonPornography)
	return x.Buf
}

func (e TL_inputReportReasonOther) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputReportReasonOther)
	x.String(e.Text)
	return x.Buf
}

func (e TL_updateNewStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateNewStickerSet)
	x.Bytes(e.Stickerset.Encode())
	return x.Buf
}

func (e TL_updateStickerSetsOrder) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateStickerSetsOrder)
	x.Int(e.Flags)
	x.VectorLong(e.Order)
	return x.Buf
}

func (e TL_updateStickerSets) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateStickerSets)
	return x.Buf
}

func (e TL_help_termsOfService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_help_termsOfService)
	x.String(e.Text)
	return x.Buf
}

func (e TL_foundGif) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_foundGif)
	x.String(e.Url)
	x.String(e.Thumb_url)
	x.String(e.Content_url)
	x.String(e.Content_type)
	x.Int(e.W)
	x.Int(e.H)
	return x.Buf
}

func (e TL_inputMediaGifExternal) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMediaGifExternal)
	x.String(e.Url)
	x.String(e.Q)
	return x.Buf
}

func (e TL_messages_foundGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_foundGifs)
	x.Int(e.Next_offset)
	x.Vector(e.Results)
	return x.Buf
}

func (e TL_inputMessagesFilterGif) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMessagesFilterGif)
	return x.Buf
}

func (e TL_updateSavedGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateSavedGifs)
	return x.Buf
}

func (e TL_updateBotInlineQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateBotInlineQuery)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.Int(e.User_id)
	x.String(e.Query)
	x.Bytes(e.Geo.Encode())
	x.String(e.Offset)
	return x.Buf
}

func (e TL_foundGifCached) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_foundGifCached)
	x.String(e.Url)
	x.Bytes(e.Photo.Encode())
	x.Bytes(e.Document.Encode())
	return x.Buf
}

func (e TL_messages_savedGifsNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_savedGifsNotModified)
	return x.Buf
}

func (e TL_messages_savedGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_savedGifs)
	x.Int(e.Hash)
	x.Vector(e.Gifs)
	return x.Buf
}

func (e TL_inputBotInlineMessageMediaAuto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputBotInlineMessageMediaAuto)
	x.Int(e.Flags)
	x.String(e.Caption)
	x.Bytes(e.Reply_markup.Encode())
	return x.Buf
}

func (e TL_inputBotInlineMessageText) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputBotInlineMessageText)
	x.Int(e.Flags)
	x.String(e.Message)
	x.Vector(e.Entities)
	x.Bytes(e.Reply_markup.Encode())
	return x.Buf
}

func (e TL_inputBotInlineResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputBotInlineResult)
	x.Int(e.Flags)
	x.String(e.Id)
	x.String(e._Type)
	x.String(e.Title)
	x.String(e.Description)
	x.String(e.Url)
	x.String(e.Thumb_url)
	x.String(e.Content_url)
	x.String(e.Content_type)
	x.Int(e.W)
	x.Int(e.H)
	x.Int(e.Duration)
	x.Bytes(e.Send_message.Encode())
	return x.Buf
}

func (e TL_botInlineMessageMediaAuto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_botInlineMessageMediaAuto)
	x.Int(e.Flags)
	x.String(e.Caption)
	x.Bytes(e.Reply_markup.Encode())
	return x.Buf
}

func (e TL_botInlineMessageText) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_botInlineMessageText)
	x.Int(e.Flags)
	x.String(e.Message)
	x.Vector(e.Entities)
	x.Bytes(e.Reply_markup.Encode())
	return x.Buf
}

func (e TL_botInlineResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_botInlineResult)
	x.Int(e.Flags)
	x.String(e.Id)
	x.String(e._Type)
	x.String(e.Title)
	x.String(e.Description)
	x.String(e.Url)
	x.String(e.Thumb_url)
	x.String(e.Content_url)
	x.String(e.Content_type)
	x.Int(e.W)
	x.Int(e.H)
	x.Int(e.Duration)
	x.Bytes(e.Send_message.Encode())
	return x.Buf
}

func (e TL_messages_botResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_botResults)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.String(e.Next_offset)
	x.Bytes(e.Switch_pm.Encode())
	x.Vector(e.Results)
	x.Int(e.Cache_time)
	return x.Buf
}

func (e TL_inputMessagesFilterVoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMessagesFilterVoice)
	return x.Buf
}

func (e TL_inputMessagesFilterMusic) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMessagesFilterMusic)
	return x.Buf
}

func (e TL_updateBotInlineSend) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateBotInlineSend)
	x.Int(e.Flags)
	x.Int(e.User_id)
	x.String(e.Query)
	x.Bytes(e.Geo.Encode())
	x.String(e.Id)
	x.Bytes(e.Msg_id.Encode())
	return x.Buf
}

func (e TL_inputPrivacyKeyChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPrivacyKeyChatInvite)
	return x.Buf
}

func (e TL_privacyKeyChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_privacyKeyChatInvite)
	return x.Buf
}

func (e TL_updateEditChannelMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateEditChannelMessage)
	x.Bytes(e.Message.Encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.Buf
}

func (e TL_exportedMessageLink) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_exportedMessageLink)
	x.String(e.Link)
	return x.Buf
}

func (e TL_messageFwdHeader) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageFwdHeader)
	x.Int(e.Flags)
	x.Int(e.From_id)
	x.Int(e.Date)
	x.Int(e.Channel_id)
	x.Int(e.Channel_post)
	x.String(e.Post_author)
	return x.Buf
}

func (e TL_messageActionPinMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionPinMessage)
	return x.Buf
}

func (e TL_peerSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_peerSettings)
	x.Int(e.Flags)
	return x.Buf
}

func (e TL_updateChannelPinnedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateChannelPinnedMessage)
	x.Int(e.Channel_id)
	x.Int(e.Id)
	return x.Buf
}

func (e TL_keyboardButtonUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_keyboardButtonUrl)
	x.String(e.Text)
	x.String(e.Url)
	return x.Buf
}

func (e TL_keyboardButtonCallback) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_keyboardButtonCallback)
	x.String(e.Text)
	x.StringBytes(e.Data)
	return x.Buf
}

func (e TL_keyboardButtonRequestPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_keyboardButtonRequestPhone)
	x.String(e.Text)
	return x.Buf
}

func (e TL_keyboardButtonRequestGeoLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_keyboardButtonRequestGeoLocation)
	x.String(e.Text)
	return x.Buf
}

func (e TL_auth_codeTypeSms) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_codeTypeSms)
	return x.Buf
}

func (e TL_auth_codeTypeCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_codeTypeCall)
	return x.Buf
}

func (e TL_auth_codeTypeFlashCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_codeTypeFlashCall)
	return x.Buf
}

func (e TL_auth_sentCodeTypeApp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_sentCodeTypeApp)
	x.Int(e.Length)
	return x.Buf
}

func (e TL_auth_sentCodeTypeSms) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_sentCodeTypeSms)
	x.Int(e.Length)
	return x.Buf
}

func (e TL_auth_sentCodeTypeCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_sentCodeTypeCall)
	x.Int(e.Length)
	return x.Buf
}

func (e TL_auth_sentCodeTypeFlashCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_sentCodeTypeFlashCall)
	x.String(e.Pattern)
	return x.Buf
}

func (e TL_keyboardButtonSwitchInline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_keyboardButtonSwitchInline)
	x.Int(e.Flags)
	x.String(e.Text)
	x.String(e.Query)
	return x.Buf
}

func (e TL_replyInlineMarkup) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_replyInlineMarkup)
	x.Vector(e.Rows)
	return x.Buf
}

func (e TL_messages_botCallbackAnswer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_botCallbackAnswer)
	x.Int(e.Flags)
	x.String(e.Message)
	x.String(e.Url)
	x.Int(e.Cache_time)
	return x.Buf
}

func (e TL_updateBotCallbackQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateBotCallbackQuery)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.Int(e.User_id)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Msg_id)
	x.Long(e.Chat_instance)
	x.StringBytes(e.Data)
	x.String(e.Game_short_name)
	return x.Buf
}

func (e TL_messages_messageEditData) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_messageEditData)
	x.Int(e.Flags)
	return x.Buf
}

func (e TL_updateEditMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateEditMessage)
	x.Bytes(e.Message.Encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.Buf
}

func (e TL_inputBotInlineMessageMediaGeo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputBotInlineMessageMediaGeo)
	x.Int(e.Flags)
	x.Bytes(e.Geo_point.Encode())
	x.Bytes(e.Reply_markup.Encode())
	return x.Buf
}

func (e TL_inputBotInlineMessageMediaVenue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputBotInlineMessageMediaVenue)
	x.Int(e.Flags)
	x.Bytes(e.Geo_point.Encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.Venue_id)
	x.Bytes(e.Reply_markup.Encode())
	return x.Buf
}

func (e TL_inputBotInlineMessageMediaContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputBotInlineMessageMediaContact)
	x.Int(e.Flags)
	x.String(e.Phone_number)
	x.String(e.First_name)
	x.String(e.Last_name)
	x.Bytes(e.Reply_markup.Encode())
	return x.Buf
}

func (e TL_botInlineMessageMediaGeo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_botInlineMessageMediaGeo)
	x.Int(e.Flags)
	x.Bytes(e.Geo.Encode())
	x.Bytes(e.Reply_markup.Encode())
	return x.Buf
}

func (e TL_botInlineMessageMediaVenue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_botInlineMessageMediaVenue)
	x.Int(e.Flags)
	x.Bytes(e.Geo.Encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.Venue_id)
	x.Bytes(e.Reply_markup.Encode())
	return x.Buf
}

func (e TL_botInlineMessageMediaContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_botInlineMessageMediaContact)
	x.Int(e.Flags)
	x.String(e.Phone_number)
	x.String(e.First_name)
	x.String(e.Last_name)
	x.Bytes(e.Reply_markup.Encode())
	return x.Buf
}

func (e TL_inputBotInlineResultPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputBotInlineResultPhoto)
	x.String(e.Id)
	x.String(e._Type)
	x.Bytes(e.Photo.Encode())
	x.Bytes(e.Send_message.Encode())
	return x.Buf
}

func (e TL_inputBotInlineResultDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputBotInlineResultDocument)
	x.Int(e.Flags)
	x.String(e.Id)
	x.String(e._Type)
	x.String(e.Title)
	x.String(e.Description)
	x.Bytes(e.Document.Encode())
	x.Bytes(e.Send_message.Encode())
	return x.Buf
}

func (e TL_botInlineMediaResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_botInlineMediaResult)
	x.Int(e.Flags)
	x.String(e.Id)
	x.String(e._Type)
	x.Bytes(e.Photo.Encode())
	x.Bytes(e.Document.Encode())
	x.String(e.Title)
	x.String(e.Description)
	x.Bytes(e.Send_message.Encode())
	return x.Buf
}

func (e TL_inputBotInlineMessageID) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputBotInlineMessageID)
	x.Int(e.Dc_id)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.Buf
}

func (e TL_updateInlineBotCallbackQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateInlineBotCallbackQuery)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.Int(e.User_id)
	x.Bytes(e.Msg_id.Encode())
	x.Long(e.Chat_instance)
	x.StringBytes(e.Data)
	x.String(e.Game_short_name)
	return x.Buf
}

func (e TL_inlineBotSwitchPM) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inlineBotSwitchPM)
	x.String(e.Text)
	x.String(e.Start_param)
	return x.Buf
}

func (e TL_messageEntityMentionName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageEntityMentionName)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.Int(e.User_id)
	return x.Buf
}

func (e TL_inputMessageEntityMentionName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMessageEntityMentionName)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.Bytes(e.User_id.Encode())
	return x.Buf
}

func (e TL_messages_peerDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_peerDialogs)
	x.Vector(e.Dialogs)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Bytes(e.State.Encode())
	return x.Buf
}

func (e TL_topPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_topPeer)
	x.Bytes(e.Peer.Encode())
	x.Double(e.Rating)
	return x.Buf
}

func (e TL_topPeerCategoryBotsPM) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_topPeerCategoryBotsPM)
	return x.Buf
}

func (e TL_topPeerCategoryBotsInline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_topPeerCategoryBotsInline)
	return x.Buf
}

func (e TL_topPeerCategoryCorrespondents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_topPeerCategoryCorrespondents)
	return x.Buf
}

func (e TL_topPeerCategoryGroups) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_topPeerCategoryGroups)
	return x.Buf
}

func (e TL_topPeerCategoryChannels) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_topPeerCategoryChannels)
	return x.Buf
}

func (e TL_topPeerCategoryPeers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_topPeerCategoryPeers)
	x.Bytes(e.Category.Encode())
	x.Int(e.Count)
	x.Vector(e.Peers)
	return x.Buf
}

func (e TL_contacts_topPeersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_topPeersNotModified)
	return x.Buf
}

func (e TL_contacts_topPeers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_topPeers)
	x.Vector(e.Categories)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_inputMessagesFilterChatPhotos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMessagesFilterChatPhotos)
	return x.Buf
}

func (e TL_updateReadChannelOutbox) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateReadChannelOutbox)
	x.Int(e.Channel_id)
	x.Int(e.Max_id)
	return x.Buf
}

func (e TL_updateDraftMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateDraftMessage)
	x.Bytes(e.Peer.Encode())
	x.Bytes(e.Draft.Encode())
	return x.Buf
}

func (e TL_draftMessageEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_draftMessageEmpty)
	return x.Buf
}

func (e TL_draftMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_draftMessage)
	x.Int(e.Flags)
	x.Int(e.Reply_to_msg_id)
	x.String(e.Message)
	x.Vector(e.Entities)
	x.Int(e.Date)
	return x.Buf
}

func (e TL_messageActionHistoryClear) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionHistoryClear)
	return x.Buf
}

func (e TL_updateReadFeaturedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateReadFeaturedStickers)
	return x.Buf
}

func (e TL_updateRecentStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateRecentStickers)
	return x.Buf
}

func (e TL_messages_featuredStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_featuredStickersNotModified)
	return x.Buf
}

func (e TL_messages_featuredStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_featuredStickers)
	x.Int(e.Hash)
	x.Vector(e.Sets)
	x.VectorLong(e.Unread)
	return x.Buf
}

func (e TL_messages_recentStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_recentStickersNotModified)
	return x.Buf
}

func (e TL_messages_recentStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_recentStickers)
	x.Int(e.Hash)
	x.Vector(e.Stickers)
	return x.Buf
}

func (e TL_messages_archivedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_archivedStickers)
	x.Int(e.Count)
	x.Vector(e.Sets)
	return x.Buf
}

func (e TL_messages_stickerSetInstallResultSuccess) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_stickerSetInstallResultSuccess)
	return x.Buf
}

func (e TL_messages_stickerSetInstallResultArchive) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_stickerSetInstallResultArchive)
	x.Vector(e.Sets)
	return x.Buf
}

func (e TL_stickerSetCovered) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_stickerSetCovered)
	x.Bytes(e.Set.Encode())
	x.Bytes(e.Cover.Encode())
	return x.Buf
}

func (e TL_inputMediaPhotoExternal) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMediaPhotoExternal)
	x.Int(e.Flags)
	x.String(e.Url)
	x.String(e.Caption)
	x.Int(e.Ttl_seconds)
	return x.Buf
}

func (e TL_inputMediaDocumentExternal) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMediaDocumentExternal)
	x.Int(e.Flags)
	x.String(e.Url)
	x.String(e.Caption)
	x.Int(e.Ttl_seconds)
	return x.Buf
}

func (e TL_updateConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateConfig)
	return x.Buf
}

func (e TL_updatePtsChanged) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updatePtsChanged)
	return x.Buf
}

func (e TL_messageActionGameScore) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionGameScore)
	x.Long(e.Game_id)
	x.Int(e.Score)
	return x.Buf
}

func (e TL_documentAttributeHasStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_documentAttributeHasStickers)
	return x.Buf
}

func (e TL_keyboardButtonGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_keyboardButtonGame)
	x.String(e.Text)
	return x.Buf
}

func (e TL_stickerSetMultiCovered) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_stickerSetMultiCovered)
	x.Bytes(e.Set.Encode())
	x.Vector(e.Covers)
	return x.Buf
}

func (e TL_maskCoords) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_maskCoords)
	x.Int(e.N)
	x.Double(e.X)
	x.Double(e.Y)
	x.Double(e.Zoom)
	return x.Buf
}

func (e TL_inputStickeredMediaPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputStickeredMediaPhoto)
	x.Bytes(e.Id.Encode())
	return x.Buf
}

func (e TL_inputStickeredMediaDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputStickeredMediaDocument)
	x.Bytes(e.Id.Encode())
	return x.Buf
}

func (e TL_inputMediaGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMediaGame)
	x.Bytes(e.Id.Encode())
	return x.Buf
}

func (e TL_messageMediaGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageMediaGame)
	x.Bytes(e.Game.Encode())
	return x.Buf
}

func (e TL_inputBotInlineMessageGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputBotInlineMessageGame)
	x.Int(e.Flags)
	x.Bytes(e.Reply_markup.Encode())
	return x.Buf
}

func (e TL_inputBotInlineResultGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputBotInlineResultGame)
	x.String(e.Id)
	x.String(e.Short_name)
	x.Bytes(e.Send_message.Encode())
	return x.Buf
}

func (e TL_game) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_game)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.String(e.Short_name)
	x.String(e.Title)
	x.String(e.Description)
	x.Bytes(e.Photo.Encode())
	x.Bytes(e.Document.Encode())
	return x.Buf
}

func (e TL_inputGameID) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputGameID)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.Buf
}

func (e TL_inputGameShortName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputGameShortName)
	x.Bytes(e.Bot_id.Encode())
	x.String(e.Short_name)
	return x.Buf
}

func (e TL_highScore) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_highScore)
	x.Int(e.Pos)
	x.Int(e.User_id)
	x.Int(e.Score)
	return x.Buf
}

func (e TL_messages_highScores) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_highScores)
	x.Vector(e.Scores)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_messages_chatsSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_chatsSlice)
	x.Int(e.Count)
	x.Vector(e.Chats)
	return x.Buf
}

func (e TL_updateChannelWebPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateChannelWebPage)
	x.Int(e.Channel_id)
	x.Bytes(e.Webpage.Encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.Buf
}

func (e TL_updates_differenceTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updates_differenceTooLong)
	x.Int(e.Pts)
	return x.Buf
}

func (e TL_sendMessageGamePlayAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_sendMessageGamePlayAction)
	return x.Buf
}

func (e TL_webPageNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_webPageNotModified)
	return x.Buf
}

func (e TL_textEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_textEmpty)
	return x.Buf
}

func (e TL_textPlain) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_textPlain)
	x.String(e.Text)
	return x.Buf
}

func (e TL_textBold) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_textBold)
	x.Bytes(e.Text.Encode())
	return x.Buf
}

func (e TL_textItalic) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_textItalic)
	x.Bytes(e.Text.Encode())
	return x.Buf
}

func (e TL_textUnderline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_textUnderline)
	x.Bytes(e.Text.Encode())
	return x.Buf
}

func (e TL_textStrike) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_textStrike)
	x.Bytes(e.Text.Encode())
	return x.Buf
}

func (e TL_textFixed) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_textFixed)
	x.Bytes(e.Text.Encode())
	return x.Buf
}

func (e TL_textUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_textUrl)
	x.Bytes(e.Text.Encode())
	x.String(e.Url)
	x.Long(e.Webpage_id)
	return x.Buf
}

func (e TL_textEmail) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_textEmail)
	x.Bytes(e.Text.Encode())
	x.String(e.Email)
	return x.Buf
}

func (e TL_textConcat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_textConcat)
	x.Vector(e.Texts)
	return x.Buf
}

func (e TL_pageBlockTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockTitle)
	x.Bytes(e.Text.Encode())
	return x.Buf
}

func (e TL_pageBlockSubtitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockSubtitle)
	x.Bytes(e.Text.Encode())
	return x.Buf
}

func (e TL_pageBlockAuthorDate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockAuthorDate)
	x.Bytes(e.Author.Encode())
	x.Int(e.Published_date)
	return x.Buf
}

func (e TL_pageBlockHeader) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockHeader)
	x.Bytes(e.Text.Encode())
	return x.Buf
}

func (e TL_pageBlockSubheader) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockSubheader)
	x.Bytes(e.Text.Encode())
	return x.Buf
}

func (e TL_pageBlockParagraph) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockParagraph)
	x.Bytes(e.Text.Encode())
	return x.Buf
}

func (e TL_pageBlockPreformatted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockPreformatted)
	x.Bytes(e.Text.Encode())
	x.String(e.Language)
	return x.Buf
}

func (e TL_pageBlockFooter) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockFooter)
	x.Bytes(e.Text.Encode())
	return x.Buf
}

func (e TL_pageBlockDivider) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockDivider)
	return x.Buf
}

func (e TL_pageBlockList) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockList)
	x.Bytes(e.Ordered.Encode())
	x.Vector(e.Items)
	return x.Buf
}

func (e TL_pageBlockBlockquote) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockBlockquote)
	x.Bytes(e.Text.Encode())
	x.Bytes(e.Caption.Encode())
	return x.Buf
}

func (e TL_pageBlockPullquote) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockPullquote)
	x.Bytes(e.Text.Encode())
	x.Bytes(e.Caption.Encode())
	return x.Buf
}

func (e TL_pageBlockPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockPhoto)
	x.Long(e.Photo_id)
	x.Bytes(e.Caption.Encode())
	return x.Buf
}

func (e TL_pageBlockVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockVideo)
	x.Int(e.Flags)
	x.Long(e.Video_id)
	x.Bytes(e.Caption.Encode())
	return x.Buf
}

func (e TL_pageBlockCover) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockCover)
	x.Bytes(e.Cover.Encode())
	return x.Buf
}

func (e TL_pageBlockEmbed) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockEmbed)
	x.Int(e.Flags)
	x.String(e.Url)
	x.String(e.Html)
	x.Long(e.Poster_photo_id)
	x.Int(e.W)
	x.Int(e.H)
	x.Bytes(e.Caption.Encode())
	return x.Buf
}

func (e TL_pageBlockEmbedPost) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockEmbedPost)
	x.String(e.Url)
	x.Long(e.Webpage_id)
	x.Long(e.Author_photo_id)
	x.String(e.Author)
	x.Int(e.Date)
	x.Vector(e.Blocks)
	x.Bytes(e.Caption.Encode())
	return x.Buf
}

func (e TL_pageBlockSlideshow) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockSlideshow)
	x.Vector(e.Items)
	x.Bytes(e.Caption.Encode())
	return x.Buf
}

func (e TL_pagePart) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pagePart)
	x.Vector(e.Blocks)
	x.Vector(e.Photos)
	x.Vector(e.Documents)
	return x.Buf
}

func (e TL_pageFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageFull)
	x.Vector(e.Blocks)
	x.Vector(e.Photos)
	x.Vector(e.Documents)
	return x.Buf
}

func (e TL_updatePhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updatePhoneCall)
	x.Bytes(e.Phone_call.Encode())
	return x.Buf
}

func (e TL_updateDialogPinned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateDialogPinned)
	x.Int(e.Flags)
	x.Bytes(e.Peer.Encode())
	return x.Buf
}

func (e TL_updatePinnedDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updatePinnedDialogs)
	x.Int(e.Flags)
	x.Vector(e.Order)
	return x.Buf
}

func (e TL_inputPrivacyKeyPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPrivacyKeyPhoneCall)
	return x.Buf
}

func (e TL_privacyKeyPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_privacyKeyPhoneCall)
	return x.Buf
}

func (e TL_pageBlockUnsupported) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockUnsupported)
	return x.Buf
}

func (e TL_pageBlockAnchor) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockAnchor)
	x.String(e.Name)
	return x.Buf
}

func (e TL_pageBlockCollage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockCollage)
	x.Vector(e.Items)
	x.Bytes(e.Caption.Encode())
	return x.Buf
}

func (e TL_inputPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPhoneCall)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.Buf
}

func (e TL_phoneCallEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phoneCallEmpty)
	x.Long(e.Id)
	return x.Buf
}

func (e TL_phoneCallWaiting) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phoneCallWaiting)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.Bytes(e.Protocol.Encode())
	x.Int(e.Receive_date)
	return x.Buf
}

func (e TL_phoneCallRequested) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phoneCallRequested)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.StringBytes(e.G_a_hash)
	x.Bytes(e.Protocol.Encode())
	return x.Buf
}

func (e TL_phoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phoneCall)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.StringBytes(e.G_a_or_b)
	x.Long(e.Key_fingerprint)
	x.Bytes(e.Protocol.Encode())
	x.Bytes(e.Connection.Encode())
	x.Vector(e.Alternative_connections)
	x.Int(e.Start_date)
	return x.Buf
}

func (e TL_phoneCallDiscarded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phoneCallDiscarded)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.Bytes(e.Reason.Encode())
	x.Int(e.Duration)
	return x.Buf
}

func (e TL_phoneConnection) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phoneConnection)
	x.Long(e.Id)
	x.String(e.Ip)
	x.String(e.Ipv6)
	x.Int(e.Port)
	x.StringBytes(e.Peer_tag)
	return x.Buf
}

func (e TL_phoneCallProtocol) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phoneCallProtocol)
	x.Int(e.Flags)
	x.Int(e.Min_layer)
	x.Int(e.Max_layer)
	return x.Buf
}

func (e TL_phone_phoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phone_phoneCall)
	x.Bytes(e.Phone_call.Encode())
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_phoneCallDiscardReasonMissed) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phoneCallDiscardReasonMissed)
	return x.Buf
}

func (e TL_phoneCallDiscardReasonDisconnect) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phoneCallDiscardReasonDisconnect)
	return x.Buf
}

func (e TL_phoneCallDiscardReasonHangup) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phoneCallDiscardReasonHangup)
	return x.Buf
}

func (e TL_phoneCallDiscardReasonBusy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phoneCallDiscardReasonBusy)
	return x.Buf
}

func (e TL_inputMessagesFilterPhoneCalls) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMessagesFilterPhoneCalls)
	x.Int(e.Flags)
	return x.Buf
}

func (e TL_messageActionPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionPhoneCall)
	x.Int(e.Flags)
	x.Long(e.Call_id)
	x.Bytes(e.Reason.Encode())
	x.Int(e.Duration)
	return x.Buf
}

func (e TL_invoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_invoice)
	x.Int(e.Flags)
	x.String(e.Currency)
	x.Vector(e.Prices)
	return x.Buf
}

func (e TL_inputMediaInvoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMediaInvoice)
	x.Int(e.Flags)
	x.String(e.Title)
	x.String(e.Description)
	x.Bytes(e.Photo.Encode())
	x.Bytes(e.Invoice.Encode())
	x.StringBytes(e.Payload)
	x.String(e.Provider)
	x.String(e.Start_param)
	return x.Buf
}

func (e TL_messageActionPaymentSentMe) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionPaymentSentMe)
	x.Int(e.Flags)
	x.String(e.Currency)
	x.Long(e.Total_amount)
	x.StringBytes(e.Payload)
	x.Bytes(e.Info.Encode())
	x.String(e.Shipping_option_id)
	x.Bytes(e.Charge.Encode())
	return x.Buf
}

func (e TL_messageMediaInvoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageMediaInvoice)
	x.Int(e.Flags)
	x.String(e.Title)
	x.String(e.Description)
	x.Bytes(e.Photo.Encode())
	x.Int(e.Receipt_msg_id)
	x.String(e.Currency)
	x.Long(e.Total_amount)
	x.String(e.Start_param)
	return x.Buf
}

func (e TL_keyboardButtonBuy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_keyboardButtonBuy)
	x.String(e.Text)
	return x.Buf
}

func (e TL_messageActionPaymentSent) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionPaymentSent)
	x.String(e.Currency)
	x.Long(e.Total_amount)
	return x.Buf
}

func (e TL_payments_paymentForm) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_payments_paymentForm)
	x.Int(e.Flags)
	x.Int(e.Bot_id)
	x.Bytes(e.Invoice.Encode())
	x.Int(e.Provider_id)
	x.String(e.Url)
	x.String(e.Native_provider)
	x.Bytes(e.Native_params.Encode())
	x.Bytes(e.Saved_info.Encode())
	x.Bytes(e.Saved_credentials.Encode())
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_postAddress) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_postAddress)
	x.String(e.Street_line1)
	x.String(e.Street_line2)
	x.String(e.City)
	x.String(e.State)
	x.String(e.Country_iso2)
	x.String(e.Post_code)
	return x.Buf
}

func (e TL_paymentRequestedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_paymentRequestedInfo)
	x.Int(e.Flags)
	x.String(e.Name)
	x.String(e.Phone)
	x.String(e.Email)
	x.Bytes(e.Shipping_address.Encode())
	return x.Buf
}

func (e TL_updateBotWebhookJSON) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateBotWebhookJSON)
	x.Bytes(e.Data.Encode())
	return x.Buf
}

func (e TL_updateBotWebhookJSONQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateBotWebhookJSONQuery)
	x.Long(e.Query_id)
	x.Bytes(e.Data.Encode())
	x.Int(e.Timeout)
	return x.Buf
}

func (e TL_updateBotShippingQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateBotShippingQuery)
	x.Long(e.Query_id)
	x.Int(e.User_id)
	x.StringBytes(e.Payload)
	x.Bytes(e.Shipping_address.Encode())
	return x.Buf
}

func (e TL_updateBotPrecheckoutQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateBotPrecheckoutQuery)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.Int(e.User_id)
	x.StringBytes(e.Payload)
	x.Bytes(e.Info.Encode())
	x.String(e.Shipping_option_id)
	x.String(e.Currency)
	x.Long(e.Total_amount)
	return x.Buf
}

func (e TL_dataJSON) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_dataJSON)
	x.String(e.Data)
	return x.Buf
}

func (e TL_labeledPrice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_labeledPrice)
	x.String(e.Label)
	x.Long(e.Amount)
	return x.Buf
}

func (e TL_paymentCharge) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_paymentCharge)
	x.String(e.Id)
	x.String(e.Provider_charge_id)
	return x.Buf
}

func (e TL_paymentSavedCredentialsCard) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_paymentSavedCredentialsCard)
	x.String(e.Id)
	x.String(e.Title)
	return x.Buf
}

func (e TL_webDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_webDocument)
	x.String(e.Url)
	x.Long(e.Access_hash)
	x.Int(e.Size)
	x.String(e.Mime_type)
	x.Vector(e.Attributes)
	x.Int(e.Dc_id)
	return x.Buf
}

func (e TL_inputWebDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputWebDocument)
	x.String(e.Url)
	x.Int(e.Size)
	x.String(e.Mime_type)
	x.Vector(e.Attributes)
	return x.Buf
}

func (e TL_inputWebFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputWebFileLocation)
	x.String(e.Url)
	x.Long(e.Access_hash)
	return x.Buf
}

func (e TL_upload_webFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_upload_webFile)
	x.Int(e.Size)
	x.String(e.Mime_type)
	x.Bytes(e.File_type.Encode())
	x.Int(e.Mtime)
	x.StringBytes(e.Bytes)
	return x.Buf
}

func (e TL_payments_validatedRequestedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_payments_validatedRequestedInfo)
	x.Int(e.Flags)
	x.String(e.Id)
	x.Vector(e.Shipping_options)
	return x.Buf
}

func (e TL_payments_paymentResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_payments_paymentResult)
	x.Bytes(e.Updates.Encode())
	return x.Buf
}

func (e TL_payments_paymentVerficationNeeded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_payments_paymentVerficationNeeded)
	x.String(e.Url)
	return x.Buf
}

func (e TL_payments_paymentReceipt) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_payments_paymentReceipt)
	x.Int(e.Flags)
	x.Int(e.Date)
	x.Int(e.Bot_id)
	x.Bytes(e.Invoice.Encode())
	x.Int(e.Provider_id)
	x.Bytes(e.Info.Encode())
	x.Bytes(e.Shipping.Encode())
	x.String(e.Currency)
	x.Long(e.Total_amount)
	x.String(e.Credentials_title)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_payments_savedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_payments_savedInfo)
	x.Int(e.Flags)
	x.Bytes(e.Saved_info.Encode())
	return x.Buf
}

func (e TL_inputPaymentCredentialsSaved) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPaymentCredentialsSaved)
	x.String(e.Id)
	x.StringBytes(e.Tmp_password)
	return x.Buf
}

func (e TL_inputPaymentCredentials) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputPaymentCredentials)
	x.Int(e.Flags)
	x.Bytes(e.Data.Encode())
	return x.Buf
}

func (e TL_account_tmpPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_tmpPassword)
	x.StringBytes(e.Tmp_password)
	x.Int(e.Valid_until)
	return x.Buf
}

func (e TL_shippingOption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_shippingOption)
	x.String(e.Id)
	x.String(e.Title)
	x.Vector(e.Prices)
	return x.Buf
}

func (e TL_phoneCallAccepted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phoneCallAccepted)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.StringBytes(e.G_b)
	x.Bytes(e.Protocol.Encode())
	return x.Buf
}

func (e TL_inputMessagesFilterRoundVoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMessagesFilterRoundVoice)
	return x.Buf
}

func (e TL_inputMessagesFilterRoundVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMessagesFilterRoundVideo)
	return x.Buf
}

func (e TL_upload_fileCdnRedirect) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_upload_fileCdnRedirect)
	x.Int(e.Dc_id)
	x.StringBytes(e.File_token)
	x.StringBytes(e.Encryption_key)
	x.StringBytes(e.Encryption_iv)
	x.Vector(e.Cdn_file_hashes)
	return x.Buf
}

func (e TL_sendMessageRecordRoundAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_sendMessageRecordRoundAction)
	return x.Buf
}

func (e TL_sendMessageUploadRoundAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_sendMessageUploadRoundAction)
	x.Int(e.Progress)
	return x.Buf
}

func (e TL_upload_cdnFileReuploadNeeded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_upload_cdnFileReuploadNeeded)
	x.StringBytes(e.Request_token)
	return x.Buf
}

func (e TL_upload_cdnFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_upload_cdnFile)
	x.StringBytes(e.Bytes)
	return x.Buf
}

func (e TL_cdnPublicKey) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_cdnPublicKey)
	x.Int(e.Dc_id)
	x.String(e.Public_key)
	return x.Buf
}

func (e TL_cdnConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_cdnConfig)
	x.Vector(e.Public_keys)
	return x.Buf
}

func (e TL_updateLangPackTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateLangPackTooLong)
	return x.Buf
}

func (e TL_updateLangPack) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateLangPack)
	x.Bytes(e.Difference.Encode())
	return x.Buf
}

func (e TL_pageBlockChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockChannel)
	x.Bytes(e.Channel.Encode())
	return x.Buf
}

func (e TL_inputStickerSetItem) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputStickerSetItem)
	x.Int(e.Flags)
	x.Bytes(e.Document.Encode())
	x.String(e.Emoji)
	x.Bytes(e.Mask_coords.Encode())
	return x.Buf
}

func (e TL_langPackString) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_langPackString)
	x.String(e.Key)
	x.String(e.Value)
	return x.Buf
}

func (e TL_langPackStringPluralized) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_langPackStringPluralized)
	x.Int(e.Flags)
	x.String(e.Key)
	x.String(e.Zero_value)
	x.String(e.One_value)
	x.String(e.Two_value)
	x.String(e.Few_value)
	x.String(e.Many_value)
	x.String(e.Other_value)
	return x.Buf
}

func (e TL_langPackStringDeleted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_langPackStringDeleted)
	x.String(e.Key)
	return x.Buf
}

func (e TL_langPackDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_langPackDifference)
	x.String(e.Lang_code)
	x.Int(e.From_version)
	x.Int(e.Version)
	x.Vector(e.Strings)
	return x.Buf
}

func (e TL_langPackLanguage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_langPackLanguage)
	x.String(e.Name)
	x.String(e.Native_name)
	x.String(e.Lang_code)
	return x.Buf
}

func (e TL_channelParticipantAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelParticipantAdmin)
	x.Int(e.Flags)
	x.Int(e.User_id)
	x.Int(e.Inviter_id)
	x.Int(e.Promoted_by)
	x.Int(e.Date)
	x.Bytes(e.Admin_rights.Encode())
	return x.Buf
}

func (e TL_channelParticipantBanned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelParticipantBanned)
	x.Int(e.Flags)
	x.Int(e.User_id)
	x.Int(e.Kicked_by)
	x.Int(e.Date)
	x.Bytes(e.Banned_rights.Encode())
	return x.Buf
}

func (e TL_channelParticipantsBanned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelParticipantsBanned)
	x.String(e.Q)
	return x.Buf
}

func (e TL_channelParticipantsSearch) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelParticipantsSearch)
	x.String(e.Q)
	return x.Buf
}

func (e TL_topPeerCategoryPhoneCalls) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_topPeerCategoryPhoneCalls)
	return x.Buf
}

func (e TL_pageBlockAudio) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_pageBlockAudio)
	x.Long(e.Audio_id)
	x.Bytes(e.Caption.Encode())
	return x.Buf
}

func (e TL_channelAdminRights) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminRights)
	x.Int(e.Flags)
	return x.Buf
}

func (e TL_channelBannedRights) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelBannedRights)
	x.Int(e.Flags)
	x.Int(e.Until_date)
	return x.Buf
}

func (e TL_channelAdminLogEventActionChangeTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminLogEventActionChangeTitle)
	x.String(e.Prev_value)
	x.String(e.New_value)
	return x.Buf
}

func (e TL_channelAdminLogEventActionChangeAbout) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminLogEventActionChangeAbout)
	x.String(e.Prev_value)
	x.String(e.New_value)
	return x.Buf
}

func (e TL_channelAdminLogEventActionChangeUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminLogEventActionChangeUsername)
	x.String(e.Prev_value)
	x.String(e.New_value)
	return x.Buf
}

func (e TL_channelAdminLogEventActionChangePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminLogEventActionChangePhoto)
	x.Bytes(e.Prev_photo.Encode())
	x.Bytes(e.New_photo.Encode())
	return x.Buf
}

func (e TL_channelAdminLogEventActionToggleInvites) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminLogEventActionToggleInvites)
	x.Bytes(e.New_value.Encode())
	return x.Buf
}

func (e TL_channelAdminLogEventActionToggleSignatures) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminLogEventActionToggleSignatures)
	x.Bytes(e.New_value.Encode())
	return x.Buf
}

func (e TL_channelAdminLogEventActionUpdatePinned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminLogEventActionUpdatePinned)
	x.Bytes(e.Message.Encode())
	return x.Buf
}

func (e TL_channelAdminLogEventActionEditMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminLogEventActionEditMessage)
	x.Bytes(e.Prev_message.Encode())
	x.Bytes(e.New_message.Encode())
	return x.Buf
}

func (e TL_channelAdminLogEventActionDeleteMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminLogEventActionDeleteMessage)
	x.Bytes(e.Message.Encode())
	return x.Buf
}

func (e TL_channelAdminLogEventActionParticipantJoin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminLogEventActionParticipantJoin)
	return x.Buf
}

func (e TL_channelAdminLogEventActionParticipantLeave) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminLogEventActionParticipantLeave)
	return x.Buf
}

func (e TL_channelAdminLogEventActionParticipantInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminLogEventActionParticipantInvite)
	x.Bytes(e.Participant.Encode())
	return x.Buf
}

func (e TL_channelAdminLogEventActionParticipantToggleBan) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminLogEventActionParticipantToggleBan)
	x.Bytes(e.Prev_participant.Encode())
	x.Bytes(e.New_participant.Encode())
	return x.Buf
}

func (e TL_channelAdminLogEventActionParticipantToggleAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminLogEventActionParticipantToggleAdmin)
	x.Bytes(e.Prev_participant.Encode())
	x.Bytes(e.New_participant.Encode())
	return x.Buf
}

func (e TL_channelAdminLogEvent) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminLogEvent)
	x.Long(e.Id)
	x.Int(e.Date)
	x.Int(e.User_id)
	x.Bytes(e.Action.Encode())
	return x.Buf
}

func (e TL_channels_adminLogResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_adminLogResults)
	x.Vector(e.Events)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_channelAdminLogEventsFilter) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminLogEventsFilter)
	x.Int(e.Flags)
	return x.Buf
}

func (e TL_messageActionScreenshotTaken) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messageActionScreenshotTaken)
	return x.Buf
}

func (e TL_popularContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_popularContact)
	x.Long(e.Client_id)
	x.Int(e.Importers)
	return x.Buf
}

func (e TL_cdnFileHash) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_cdnFileHash)
	x.Int(e.Offset)
	x.Int(e.Limit)
	x.StringBytes(e.Hash)
	return x.Buf
}

func (e TL_inputMessagesFilterMyMentions) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMessagesFilterMyMentions)
	return x.Buf
}

func (e TL_inputMessagesFilterMyMentionsUnread) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_inputMessagesFilterMyMentionsUnread)
	return x.Buf
}

func (e TL_updateContactsReset) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateContactsReset)
	return x.Buf
}

func (e TL_channelAdminLogEventActionChangeStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channelAdminLogEventActionChangeStickerSet)
	x.Bytes(e.Prev_stickerset.Encode())
	x.Bytes(e.New_stickerset.Encode())
	return x.Buf
}

func (e TL_updateFavedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateFavedStickers)
	return x.Buf
}

func (e TL_messages_favedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_favedStickers)
	x.Int(e.Hash)
	x.Vector(e.Packs)
	x.Vector(e.Stickers)
	return x.Buf
}

func (e TL_messages_favedStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_favedStickersNotModified)
	return x.Buf
}

func (e TL_updateChannelReadMessagesContents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updateChannelReadMessagesContents)
	x.Int(e.Channel_id)
	x.VectorInt(e.Messages)
	return x.Buf
}

func (e TL_invokeAfterMsg) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_invokeAfterMsg)
	x.Long(e.Msg_id)
	x.Bytes(e.Query.Encode())
	return x.Buf
}

func (e TL_invokeAfterMsgs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_invokeAfterMsgs)
	x.VectorLong(e.Msg_ids)
	x.Bytes(e.Query.Encode())
	return x.Buf
}

func (e TL_auth_checkPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_checkPhone)
	x.String(e.PhoneNumber)
	return x.Buf
}

func (e TL_auth_sendCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_sendCode)
	x.Int(e.Flags)
	x.String(e.PhoneNumber)
	x.Bytes(e.CurrentNumber.Encode())
	x.Int(e.ApiID)
	x.String(e.ApiHash)
	return x.Buf
}

func (e TL_auth_signUp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_signUp)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	x.String(e.PhoneCode)
	x.String(e.FirstName)
	x.String(e.LastName)
	return x.Buf
}

func (e TL_auth_signIn) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_signIn)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	x.String(e.PhoneCode)
	return x.Buf
}

func (e TL_auth_logOut) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_logOut)
	return x.Buf
}

func (e TL_auth_resetAuthorizations) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_resetAuthorizations)
	return x.Buf
}

func (e TL_auth_sendInvites) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_sendInvites)
	x.VectorString(e.Phone_numbers)
	x.String(e.Message)
	return x.Buf
}

func (e TL_auth_exportAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_exportAuthorization)
	x.Int(e.Dc_id)
	return x.Buf
}

func (e TL_auth_importAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_importAuthorization)
	x.Int(e.Id)
	x.StringBytes(e.Bytes)
	return x.Buf
}

func (e TL_account_registerDevice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_registerDevice)
	x.Int(e.Token_type)
	x.String(e.Token)
	return x.Buf
}

func (e TL_account_unregisterDevice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_unregisterDevice)
	x.Int(e.Token_type)
	x.String(e.Token)
	return x.Buf
}

func (e TL_account_updateNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_updateNotifySettings)
	x.Bytes(e.Peer.Encode())
	x.Bytes(e.Settings.Encode())
	return x.Buf
}

func (e TL_account_getNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_getNotifySettings)
	x.Bytes(e.Peer.Encode())
	return x.Buf
}

func (e TL_account_resetNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_resetNotifySettings)
	return x.Buf
}

func (e TL_account_updateProfile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_updateProfile)
	x.Int(e.Flags)
	x.String(e.First_name)
	x.String(e.Last_name)
	x.String(e.About)
	return x.Buf
}

func (e TL_account_updateStatus) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_updateStatus)
	x.Bytes(e.Offline.Encode())
	return x.Buf
}

func (e TL_account_getWallPapers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_getWallPapers)
	return x.Buf
}

func (e TL_users_getUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_users_getUsers)
	x.Vector(e.Id)
	return x.Buf
}

func (e TL_users_getFullUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_users_getFullUser)
	x.Bytes(e.Id.Encode())
	return x.Buf
}

func (e TL_contacts_getStatuses) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_getStatuses)
	return x.Buf
}

func (e TL_contacts_getContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_getContacts)
	x.Int(e.Hash)
	return x.Buf
}

func (e TL_contacts_importContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_importContacts)
	x.Vector(e.Contacts)
	return x.Buf
}

func (e TL_contacts_search) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_search)
	x.String(e.Q)
	x.Int(e.Limit)
	return x.Buf
}

func (e TL_contacts_deleteContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_deleteContact)
	x.Bytes(e.Id.Encode())
	return x.Buf
}

func (e TL_contacts_deleteContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_deleteContacts)
	x.Vector(e.Id)
	return x.Buf
}

func (e TL_contacts_block) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_block)
	x.Bytes(e.Id.Encode())
	return x.Buf
}

func (e TL_contacts_unblock) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_unblock)
	x.Bytes(e.Id.Encode())
	return x.Buf
}

func (e TL_contacts_getBlocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_getBlocked)
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.Buf
}

func (e TL_messages_getMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getMessages)
	x.VectorInt(e.Id)
	return x.Buf
}

func (e TL_messages_getDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getDialogs)
	x.Int(e.Flags)
	x.Int(e.Offset_date)
	x.Int(e.Offset_id)
	x.Bytes(e.Offset_peer.Encode())
	x.Int(e.Limit)
	return x.Buf
}

func (e TL_messages_getHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getHistory)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Offset_id)
	x.Int(e.Offset_date)
	x.Int(e.Add_offset)
	x.Int(e.Limit)
	x.Int(e.Max_id)
	x.Int(e.Min_id)
	return x.Buf
}

func (e TL_messages_search) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_search)
	x.Int(e.Flags)
	x.Bytes(e.Peer.Encode())
	x.String(e.Q)
	x.Bytes(e.From_id.Encode())
	x.Bytes(e.Filter.Encode())
	x.Int(e.Min_date)
	x.Int(e.Max_date)
	x.Int(e.Offset_id)
	x.Int(e.Add_offset)
	x.Int(e.Limit)
	x.Int(e.Max_id)
	x.Int(e.Min_id)
	return x.Buf
}

func (e TL_messages_readHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_readHistory)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Max_id)
	return x.Buf
}

func (e TL_messages_deleteHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_deleteHistory)
	x.Int(e.Flags)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Max_id)
	return x.Buf
}

func (e TL_messages_deleteMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_deleteMessages)
	x.Int(e.Flags)
	x.VectorInt(e.Id)
	return x.Buf
}

func (e TL_messages_receivedMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_receivedMessages)
	x.Int(e.Max_id)
	return x.Buf
}

func (e TL_messages_setTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_setTyping)
	x.Bytes(e.Peer.Encode())
	x.Bytes(e.Action.Encode())
	return x.Buf
}

func (e TL_messages_sendMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_sendMessage)

	var flags uint32 = 0
	if e.Reply_to_msg_id != 0 {
		flags |= 1 << 0
	}
	if e.Reply_markup != nil {
		flags |= 1 << 2
	}
	if e.Entities != nil {
		flags |= 1 << 3
	}

	x.Int(e.Flags)
	x.Bytes(e.Peer.Encode())

	if e.Reply_to_msg_id != 0 {
		x.Int(e.Reply_to_msg_id)
	}
	x.String(e.Message)
	x.Long(e.Random_id)

	if e.Reply_markup != nil {
		x.Bytes(e.Reply_markup.Encode())
	}
	if e.Entities != nil {
		x.Vector(e.Entities)
	}

	return x.Buf
}

func (e TL_messages_sendMedia) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_sendMedia)
	x.Int(e.Flags)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Reply_to_msg_id)
	x.Bytes(e.Media.Encode())
	x.Long(e.Random_id)
	x.Bytes(e.Reply_markup.Encode())
	return x.Buf
}

func (e TL_messages_forwardMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_forwardMessages)
	x.Int(e.Flags)
	x.Bytes(e.From_peer.Encode())
	x.VectorInt(e.Id)
	x.VectorLong(e.Random_id)
	x.Bytes(e.To_peer.Encode())
	return x.Buf
}

func (e TL_messages_getChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getChats)
	x.VectorInt(e.Id)
	return x.Buf
}

func (e TL_messages_getFullChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getFullChat)
	x.Int(e.Chat_id)
	return x.Buf
}

func (e TL_messages_editChatTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_editChatTitle)
	x.Int(e.Chat_id)
	x.String(e.Title)
	return x.Buf
}

func (e TL_messages_editChatPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_editChatPhoto)
	x.Int(e.Chat_id)
	x.Bytes(e.Photo.Encode())
	return x.Buf
}

func (e TL_messages_addChatUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_addChatUser)
	x.Int(e.Chat_id)
	x.Bytes(e.User_id.Encode())
	x.Int(e.Fwd_limit)
	return x.Buf
}

func (e TL_messages_deleteChatUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_deleteChatUser)
	x.Int(e.Chat_id)
	x.Bytes(e.User_id.Encode())
	return x.Buf
}

func (e TL_messages_createChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_createChat)
	x.Vector(e.Users)
	x.String(e.Title)
	return x.Buf
}

func (e TL_updates_getState) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updates_getState)
	return x.Buf
}

func (e TL_updates_getDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updates_getDifference)
	x.Int(e.Flags)
	x.Int(e.Pts)
	x.Int(e.Pts_total_limit)
	x.Int(e.Date)
	x.Int(e.Qts)
	return x.Buf
}

func (e TL_photos_updateProfilePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_photos_updateProfilePhoto)
	x.Bytes(e.Id.Encode())
	return x.Buf
}

func (e TL_photos_uploadProfilePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_photos_uploadProfilePhoto)
	x.Bytes(e.File.Encode())
	return x.Buf
}

func (e TL_upload_saveFilePart) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_upload_saveFilePart)
	x.Long(e.File_id)
	x.Int(e.File_part)
	x.StringBytes(e.Bytes)
	return x.Buf
}

func (e TL_upload_getFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_upload_getFile)
	x.Bytes(e.Location.Encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.Buf
}

func (e TL_help_getConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_help_getConfig)
	return x.Buf
}

func (e TL_help_getNearestDc) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_help_getNearestDc)
	return x.Buf
}

func (e TL_help_getAppUpdate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_help_getAppUpdate)
	return x.Buf
}

func (e TL_help_saveAppLog) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_help_saveAppLog)
	x.Vector(e.Events)
	return x.Buf
}

func (e TL_help_getInviteText) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_help_getInviteText)
	return x.Buf
}

func (e TL_photos_deletePhotos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_photos_deletePhotos)
	x.Vector(e.Id)
	return x.Buf
}

func (e TL_photos_getUserPhotos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_photos_getUserPhotos)
	x.Bytes(e.User_id.Encode())
	x.Int(e.Offset)
	x.Long(e.Max_id)
	x.Int(e.Limit)
	return x.Buf
}

func (e TL_messages_forwardMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_forwardMessage)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Id)
	x.Long(e.Random_id)
	return x.Buf
}

func (e TL_messages_getDhConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getDhConfig)
	x.Int(e.Version)
	x.Int(e.Random_length)
	return x.Buf
}

func (e TL_messages_requestEncryption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_requestEncryption)
	x.Bytes(e.User_id.Encode())
	x.Int(e.Random_id)
	x.StringBytes(e.G_a)
	return x.Buf
}

func (e TL_messages_acceptEncryption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_acceptEncryption)
	x.Bytes(e.Peer.Encode())
	x.StringBytes(e.G_b)
	x.Long(e.Key_fingerprint)
	return x.Buf
}

func (e TL_messages_discardEncryption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_discardEncryption)
	x.Int(e.Chat_id)
	return x.Buf
}

func (e TL_messages_setEncryptedTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_setEncryptedTyping)
	x.Bytes(e.Peer.Encode())
	x.Bytes(e.Typing.Encode())
	return x.Buf
}

func (e TL_messages_readEncryptedHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_readEncryptedHistory)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Max_date)
	return x.Buf
}

func (e TL_messages_sendEncrypted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_sendEncrypted)
	x.Bytes(e.Peer.Encode())
	x.Long(e.Random_id)
	x.StringBytes(e.Data)
	return x.Buf
}

func (e TL_messages_sendEncryptedFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_sendEncryptedFile)
	x.Bytes(e.Peer.Encode())
	x.Long(e.Random_id)
	x.StringBytes(e.Data)
	x.Bytes(e.File.Encode())
	return x.Buf
}

func (e TL_messages_sendEncryptedService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_sendEncryptedService)
	x.Bytes(e.Peer.Encode())
	x.Long(e.Random_id)
	x.StringBytes(e.Data)
	return x.Buf
}

func (e TL_messages_receivedQueue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_receivedQueue)
	x.Int(e.Max_qts)
	return x.Buf
}

func (e TL_upload_saveBigFilePart) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_upload_saveBigFilePart)
	x.Long(e.File_id)
	x.Int(e.File_part)
	x.Int(e.File_total_parts)
	x.StringBytes(e.Bytes)
	return x.Buf
}

func (e TL_initConnection) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_initConnection)
	x.Int(e.Api_id)
	x.String(e.Device_model)
	x.String(e.System_version)
	x.String(e.App_version)
	x.String(e.System_lang_code)
	x.String(e.Lang_pack)
	x.String(e.Lang_code)
	x.Bytes(e.Query.Encode())
	return x.Buf
}

func (e TL_help_getSupport) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_help_getSupport)
	return x.Buf
}

func (e TL_auth_bindTempAuthKey) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_bindTempAuthKey)
	x.Long(e.Perm_auth_key_id)
	x.Long(e.Nonce)
	x.Int(e.Expires_at)
	x.StringBytes(e.Encrypted_message)
	return x.Buf
}

func (e TL_contacts_exportCard) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_exportCard)
	return x.Buf
}

func (e TL_contacts_importCard) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_importCard)
	x.VectorInt(e.Export_card)
	return x.Buf
}

func (e TL_messages_readMessageContents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_readMessageContents)
	x.VectorInt(e.Id)
	return x.Buf
}

func (e TL_account_checkUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_checkUsername)
	x.String(e.Username)
	return x.Buf
}

func (e TL_account_updateUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_updateUsername)
	x.String(e.Username)
	return x.Buf
}

func (e TL_account_getPrivacy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_getPrivacy)
	x.Bytes(e.Key.Encode())
	return x.Buf
}

func (e TL_account_setPrivacy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_setPrivacy)
	x.Bytes(e.Key.Encode())
	x.Vector(e.Rules)
	return x.Buf
}

func (e TL_account_deleteAccount) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_deleteAccount)
	x.String(e.Reason)
	return x.Buf
}

func (e TL_account_getAccountTTL) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_getAccountTTL)
	return x.Buf
}

func (e TL_account_setAccountTTL) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_setAccountTTL)
	x.Bytes(e.Ttl.Encode())
	return x.Buf
}

func (e TL_invokeWithLayer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_invokeWithLayer)
	x.Int(e.Layer)
	x.Bytes(e.Query.Encode())
	return x.Buf
}

func (e TL_contacts_resolveUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_resolveUsername)
	x.String(e.Username)
	return x.Buf
}

func (e TL_account_sendChangePhoneCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_sendChangePhoneCode)
	x.Int(e.Flags)
	x.String(e.Phone_number)
	x.Bytes(e.Current_number.Encode())
	return x.Buf
}

func (e TL_account_changePhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_changePhone)
	x.String(e.Phone_number)
	x.String(e.Phone_code_hash)
	x.String(e.Phone_code)
	return x.Buf
}

func (e TL_messages_getAllStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getAllStickers)
	x.Int(e.Hash)
	return x.Buf
}

func (e TL_account_updateDeviceLocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_updateDeviceLocked)
	x.Int(e.Period)
	return x.Buf
}

func (e TL_account_getPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_getPassword)
	return x.Buf
}

func (e TL_auth_checkPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_checkPassword)
	x.StringBytes(e.Password_hash)
	return x.Buf
}

func (e TL_messages_getWebPagePreview) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getWebPagePreview)
	x.String(e.Message)
	return x.Buf
}

func (e TL_account_getAuthorizations) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_getAuthorizations)
	return x.Buf
}

func (e TL_account_resetAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_resetAuthorization)
	x.Long(e.Hash)
	return x.Buf
}

func (e TL_account_getPasswordSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_getPasswordSettings)
	x.StringBytes(e.Current_password_hash)
	return x.Buf
}

func (e TL_account_updatePasswordSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_updatePasswordSettings)
	x.StringBytes(e.Current_password_hash)
	x.Bytes(e.New_settings.Encode())
	return x.Buf
}

func (e TL_auth_requestPasswordRecovery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_requestPasswordRecovery)
	return x.Buf
}

func (e TL_auth_recoverPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_recoverPassword)
	x.String(e.Code)
	return x.Buf
}

func (e TL_invokeWithoutUpdates) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_invokeWithoutUpdates)
	x.Bytes(e.Query.Encode())
	return x.Buf
}

func (e TL_messages_exportChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_exportChatInvite)
	x.Int(e.Chat_id)
	return x.Buf
}

func (e TL_messages_checkChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_checkChatInvite)
	x.String(e.Hash)
	return x.Buf
}

func (e TL_messages_importChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_importChatInvite)
	x.String(e.Hash)
	return x.Buf
}

func (e TL_messages_getStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getStickerSet)
	x.Bytes(e.Stickerset.Encode())
	return x.Buf
}

func (e TL_messages_installStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_installStickerSet)
	x.Bytes(e.Stickerset.Encode())
	x.Bytes(e.Archived.Encode())
	return x.Buf
}

func (e TL_messages_uninstallStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_uninstallStickerSet)
	x.Bytes(e.Stickerset.Encode())
	return x.Buf
}

func (e TL_auth_importBotAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_importBotAuthorization)
	x.Int(e.Flags)
	x.Int(e.Api_id)
	x.String(e.Api_hash)
	x.String(e.Bot_auth_token)
	return x.Buf
}

func (e TL_messages_startBot) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_startBot)
	x.Bytes(e.Bot.Encode())
	x.Bytes(e.Peer.Encode())
	x.Long(e.Random_id)
	x.String(e.Start_param)
	return x.Buf
}

func (e TL_help_getAppChangelog) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_help_getAppChangelog)
	x.String(e.Prev_app_version)
	return x.Buf
}

func (e TL_messages_reportSpam) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_reportSpam)
	x.Bytes(e.Peer.Encode())
	return x.Buf
}

func (e TL_messages_getMessagesViews) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getMessagesViews)
	x.Bytes(e.Peer.Encode())
	x.VectorInt(e.Id)
	x.Bytes(e.Increment.Encode())
	return x.Buf
}

func (e TL_updates_getChannelDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_updates_getChannelDifference)
	x.Int(e.Flags)
	x.Bytes(e.Channel.Encode())
	x.Bytes(e.Filter.Encode())
	x.Int(e.Pts)
	x.Int(e.Limit)
	return x.Buf
}

func (e TL_channels_readHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_readHistory)
	x.Bytes(e.Channel.Encode())
	x.Int(e.Max_id)
	return x.Buf
}

func (e TL_channels_deleteMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_deleteMessages)
	x.Bytes(e.Channel.Encode())
	x.VectorInt(e.Id)
	return x.Buf
}

func (e TL_channels_deleteUserHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_deleteUserHistory)
	x.Bytes(e.Channel.Encode())
	x.Bytes(e.User_id.Encode())
	return x.Buf
}

func (e TL_channels_reportSpam) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_reportSpam)
	x.Bytes(e.Channel.Encode())
	x.Bytes(e.User_id.Encode())
	x.VectorInt(e.Id)
	return x.Buf
}

func (e TL_channels_getMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_getMessages)
	x.Bytes(e.Channel.Encode())
	x.VectorInt(e.Id)
	return x.Buf
}

func (e TL_channels_getParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_getParticipants)
	x.Bytes(e.Channel.Encode())
	x.Bytes(e.Filter.Encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.Buf
}

func (e TL_channels_getParticipant) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_getParticipant)
	x.Bytes(e.Channel.Encode())
	x.Bytes(e.User_id.Encode())
	return x.Buf
}

func (e TL_channels_getChannels) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_getChannels)
	x.Vector(e.Id)
	return x.Buf
}

func (e TL_channels_getFullChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_getFullChannel)
	x.Bytes(e.Channel.Encode())
	return x.Buf
}

func (e TL_channels_createChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_createChannel)
	x.Int(e.Flags)
	x.String(e.Title)
	x.String(e.About)
	return x.Buf
}

func (e TL_channels_editAbout) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_editAbout)
	x.Bytes(e.Channel.Encode())
	x.String(e.About)
	return x.Buf
}

func (e TL_channels_editAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_editAdmin)
	x.Bytes(e.Channel.Encode())
	x.Bytes(e.User_id.Encode())
	x.Bytes(e.Admin_rights.Encode())
	return x.Buf
}

func (e TL_channels_editTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_editTitle)
	x.Bytes(e.Channel.Encode())
	x.String(e.Title)
	return x.Buf
}

func (e TL_channels_editPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_editPhoto)
	x.Bytes(e.Channel.Encode())
	x.Bytes(e.Photo.Encode())
	return x.Buf
}

func (e TL_channels_checkUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_checkUsername)
	x.Bytes(e.Channel.Encode())
	x.String(e.Username)
	return x.Buf
}

func (e TL_channels_updateUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_updateUsername)
	x.Bytes(e.Channel.Encode())
	x.String(e.Username)
	return x.Buf
}

func (e TL_channels_joinChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_joinChannel)
	x.Bytes(e.Channel.Encode())
	return x.Buf
}

func (e TL_channels_leaveChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_leaveChannel)
	x.Bytes(e.Channel.Encode())
	return x.Buf
}

func (e TL_channels_inviteToChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_inviteToChannel)
	x.Bytes(e.Channel.Encode())
	x.Vector(e.Users)
	return x.Buf
}

func (e TL_channels_exportInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_exportInvite)
	x.Bytes(e.Channel.Encode())
	return x.Buf
}

func (e TL_channels_deleteChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_deleteChannel)
	x.Bytes(e.Channel.Encode())
	return x.Buf
}

func (e TL_messages_toggleChatAdmins) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_toggleChatAdmins)
	x.Int(e.Chat_id)
	x.Bytes(e.Enabled.Encode())
	return x.Buf
}

func (e TL_messages_editChatAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_editChatAdmin)
	x.Int(e.Chat_id)
	x.Bytes(e.User_id.Encode())
	x.Bytes(e.Is_admin.Encode())
	return x.Buf
}

func (e TL_messages_migrateChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_migrateChat)
	x.Int(e.Chat_id)
	return x.Buf
}

func (e TL_messages_searchGlobal) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_searchGlobal)
	x.String(e.Q)
	x.Int(e.Offset_date)
	x.Bytes(e.Offset_peer.Encode())
	x.Int(e.Offset_id)
	x.Int(e.Limit)
	return x.Buf
}

func (e TL_account_reportPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_reportPeer)
	x.Bytes(e.Peer.Encode())
	x.Bytes(e.Reason.Encode())
	return x.Buf
}

func (e TL_messages_reorderStickerSets) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_reorderStickerSets)
	x.Int(e.Flags)
	x.VectorLong(e.Order)
	return x.Buf
}

func (e TL_help_getTermsOfService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_help_getTermsOfService)
	return x.Buf
}

func (e TL_messages_getDocumentByHash) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getDocumentByHash)
	x.StringBytes(e.Sha256)
	x.Int(e.Size)
	x.String(e.Mime_type)
	return x.Buf
}

func (e TL_messages_searchGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_searchGifs)
	x.String(e.Q)
	x.Int(e.Offset)
	return x.Buf
}

func (e TL_messages_getSavedGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getSavedGifs)
	x.Int(e.Hash)
	return x.Buf
}

func (e TL_messages_saveGif) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_saveGif)
	x.Bytes(e.Id.Encode())
	x.Bytes(e.Unsave.Encode())
	return x.Buf
}

func (e TL_messages_getInlineBotResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getInlineBotResults)
	x.Int(e.Flags)
	x.Bytes(e.Bot.Encode())
	x.Bytes(e.Peer.Encode())
	x.Bytes(e.Geo_point.Encode())
	x.String(e.Query)
	x.String(e.Offset)
	return x.Buf
}

func (e TL_messages_setInlineBotResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_setInlineBotResults)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.Vector(e.Results)
	x.Int(e.Cache_time)
	x.String(e.Next_offset)
	x.Bytes(e.Switch_pm.Encode())
	return x.Buf
}

func (e TL_messages_sendInlineBotResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_sendInlineBotResult)
	x.Int(e.Flags)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Reply_to_msg_id)
	x.Long(e.Random_id)
	x.Long(e.Query_id)
	x.String(e.Id)
	return x.Buf
}

func (e TL_channels_toggleInvites) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_toggleInvites)
	x.Bytes(e.Channel.Encode())
	x.Bytes(e.Enabled.Encode())
	return x.Buf
}

func (e TL_channels_exportMessageLink) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_exportMessageLink)
	x.Bytes(e.Channel.Encode())
	x.Int(e.Id)
	return x.Buf
}

func (e TL_channels_toggleSignatures) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_toggleSignatures)
	x.Bytes(e.Channel.Encode())
	x.Bytes(e.Enabled.Encode())
	return x.Buf
}

func (e TL_messages_hideReportSpam) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_hideReportSpam)
	x.Bytes(e.Peer.Encode())
	return x.Buf
}

func (e TL_messages_getPeerSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getPeerSettings)
	x.Bytes(e.Peer.Encode())
	return x.Buf
}

func (e TL_channels_updatePinnedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_updatePinnedMessage)
	x.Int(e.Flags)
	x.Bytes(e.Channel.Encode())
	x.Int(e.Id)
	return x.Buf
}

func (e TL_auth_resendCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_resendCode)
	x.String(e.Phone_number)
	x.String(e.Phone_code_hash)
	return x.Buf
}

func (e TL_auth_cancelCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_cancelCode)
	x.String(e.Phone_number)
	x.String(e.Phone_code_hash)
	return x.Buf
}

func (e TL_messages_getMessageEditData) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getMessageEditData)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Id)
	return x.Buf
}

func (e TL_messages_editMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_editMessage)
	x.Int(e.Flags)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Id)
	x.String(e.Message)
	x.Bytes(e.Reply_markup.Encode())
	x.Vector(e.Entities)
	return x.Buf
}

func (e TL_messages_editInlineBotMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_editInlineBotMessage)
	x.Int(e.Flags)
	x.Bytes(e.Id.Encode())
	x.String(e.Message)
	x.Bytes(e.Reply_markup.Encode())
	x.Vector(e.Entities)
	return x.Buf
}

func (e TL_messages_getBotCallbackAnswer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getBotCallbackAnswer)
	x.Int(e.Flags)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Msg_id)
	x.StringBytes(e.Data)
	return x.Buf
}

func (e TL_messages_setBotCallbackAnswer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_setBotCallbackAnswer)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.String(e.Message)
	x.String(e.Url)
	x.Int(e.Cache_time)
	return x.Buf
}

func (e TL_contacts_getTopPeers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_getTopPeers)
	x.Int(e.Flags)
	x.Int(e.Offset)
	x.Int(e.Limit)
	x.Int(e.Hash)
	return x.Buf
}

func (e TL_contacts_resetTopPeerRating) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_resetTopPeerRating)
	x.Bytes(e.Category.Encode())
	x.Bytes(e.Peer.Encode())
	return x.Buf
}

func (e TL_messages_getPeerDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getPeerDialogs)
	x.Vector(e.Peers)
	return x.Buf
}

func (e TL_messages_saveDraft) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_saveDraft)
	x.Int(e.Flags)
	x.Int(e.Reply_to_msg_id)
	x.Bytes(e.Peer.Encode())
	x.String(e.Message)
	x.Vector(e.Entities)
	return x.Buf
}

func (e TL_messages_getAllDrafts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getAllDrafts)
	return x.Buf
}

func (e TL_account_sendConfirmPhoneCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_sendConfirmPhoneCode)
	x.Int(e.Flags)
	x.String(e.Hash)
	x.Bytes(e.Current_number.Encode())
	return x.Buf
}

func (e TL_account_confirmPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_confirmPhone)
	x.String(e.Phone_code_hash)
	x.String(e.Phone_code)
	return x.Buf
}

func (e TL_messages_getFeaturedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getFeaturedStickers)
	x.Int(e.Hash)
	return x.Buf
}

func (e TL_messages_readFeaturedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_readFeaturedStickers)
	x.VectorLong(e.Id)
	return x.Buf
}

func (e TL_messages_getRecentStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getRecentStickers)
	x.Int(e.Flags)
	x.Int(e.Hash)
	return x.Buf
}

func (e TL_messages_saveRecentSticker) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_saveRecentSticker)
	x.Int(e.Flags)
	x.Bytes(e.Id.Encode())
	x.Bytes(e.Unsave.Encode())
	return x.Buf
}

func (e TL_messages_clearRecentStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_clearRecentStickers)
	x.Int(e.Flags)
	return x.Buf
}

func (e TL_messages_getArchivedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getArchivedStickers)
	x.Int(e.Flags)
	x.Long(e.Offset_id)
	x.Int(e.Limit)
	return x.Buf
}

func (e TL_channels_getAdminedPublicChannels) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_getAdminedPublicChannels)
	return x.Buf
}

func (e TL_auth_dropTempAuthKeys) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_auth_dropTempAuthKeys)
	x.VectorLong(e.Except_auth_keys)
	return x.Buf
}

func (e TL_messages_setGameScore) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_setGameScore)
	x.Int(e.Flags)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Id)
	x.Bytes(e.User_id.Encode())
	x.Int(e.Score)
	return x.Buf
}

func (e TL_messages_setInlineGameScore) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_setInlineGameScore)
	x.Int(e.Flags)
	x.Bytes(e.Id.Encode())
	x.Bytes(e.User_id.Encode())
	x.Int(e.Score)
	return x.Buf
}

func (e TL_messages_getMaskStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getMaskStickers)
	x.Int(e.Hash)
	return x.Buf
}

func (e TL_messages_getAttachedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getAttachedStickers)
	x.Bytes(e.Media.Encode())
	return x.Buf
}

func (e TL_messages_getGameHighScores) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getGameHighScores)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Id)
	x.Bytes(e.User_id.Encode())
	return x.Buf
}

func (e TL_messages_getInlineGameHighScores) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getInlineGameHighScores)
	x.Bytes(e.Id.Encode())
	x.Bytes(e.User_id.Encode())
	return x.Buf
}

func (e TL_messages_getCommonChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getCommonChats)
	x.Bytes(e.User_id.Encode())
	x.Int(e.Max_id)
	x.Int(e.Limit)
	return x.Buf
}

func (e TL_messages_getAllChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getAllChats)
	x.VectorInt(e.Except_ids)
	return x.Buf
}

func (e TL_help_setBotUpdatesStatus) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_help_setBotUpdatesStatus)
	x.Int(e.Pending_updates_count)
	x.String(e.Message)
	return x.Buf
}

func (e TL_messages_getWebPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getWebPage)
	x.String(e.Url)
	x.Int(e.Hash)
	return x.Buf
}

func (e TL_messages_toggleDialogPin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_toggleDialogPin)
	x.Int(e.Flags)
	x.Bytes(e.Peer.Encode())
	return x.Buf
}

func (e TL_messages_reorderPinnedDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_reorderPinnedDialogs)
	x.Int(e.Flags)
	x.Vector(e.Order)
	return x.Buf
}

func (e TL_messages_getPinnedDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getPinnedDialogs)
	return x.Buf
}

func (e TL_phone_requestCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phone_requestCall)
	x.Bytes(e.User_id.Encode())
	x.Int(e.Random_id)
	x.StringBytes(e.G_a_hash)
	x.Bytes(e.Protocol.Encode())
	return x.Buf
}

func (e TL_phone_acceptCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phone_acceptCall)
	x.Bytes(e.Peer.Encode())
	x.StringBytes(e.G_b)
	x.Bytes(e.Protocol.Encode())
	return x.Buf
}

func (e TL_phone_discardCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phone_discardCall)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Duration)
	x.Bytes(e.Reason.Encode())
	x.Long(e.Connection_id)
	return x.Buf
}

func (e TL_phone_receivedCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phone_receivedCall)
	x.Bytes(e.Peer.Encode())
	return x.Buf
}

func (e TL_messages_reportEncryptedSpam) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_reportEncryptedSpam)
	x.Bytes(e.Peer.Encode())
	return x.Buf
}

func (e TL_payments_getPaymentForm) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_payments_getPaymentForm)
	x.Int(e.Msg_id)
	return x.Buf
}

func (e TL_payments_sendPaymentForm) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_payments_sendPaymentForm)
	x.Int(e.Flags)
	x.Int(e.Msg_id)
	x.String(e.Requested_info_id)
	x.String(e.Shipping_option_id)
	x.Bytes(e.Credentials.Encode())
	return x.Buf
}

func (e TL_account_getTmpPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_account_getTmpPassword)
	x.StringBytes(e.Password_hash)
	x.Int(e.Period)
	return x.Buf
}

func (e TL_messages_setBotShippingResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_setBotShippingResults)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.String(e.Error)
	x.Vector(e.Shipping_options)
	return x.Buf
}

func (e TL_messages_setBotPrecheckoutResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_setBotPrecheckoutResults)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.String(e.Error)
	return x.Buf
}

func (e TL_upload_getWebFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_upload_getWebFile)
	x.Bytes(e.Location.Encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.Buf
}

func (e TL_bots_sendCustomRequest) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_bots_sendCustomRequest)
	x.String(e.Custom_method)
	x.Bytes(e.Params.Encode())
	return x.Buf
}

func (e TL_bots_answerWebhookJSONQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_bots_answerWebhookJSONQuery)
	x.Long(e.Query_id)
	x.Bytes(e.Data.Encode())
	return x.Buf
}

func (e TL_payments_getPaymentReceipt) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_payments_getPaymentReceipt)
	x.Int(e.Msg_id)
	return x.Buf
}

func (e TL_payments_validateRequestedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_payments_validateRequestedInfo)
	x.Int(e.Flags)
	x.Int(e.Msg_id)
	x.Bytes(e.Info.Encode())
	return x.Buf
}

func (e TL_payments_getSavedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_payments_getSavedInfo)
	return x.Buf
}

func (e TL_payments_clearSavedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_payments_clearSavedInfo)
	x.Int(e.Flags)
	return x.Buf
}

func (e TL_phone_getCallConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phone_getCallConfig)
	return x.Buf
}

func (e TL_phone_confirmCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phone_confirmCall)
	x.Bytes(e.Peer.Encode())
	x.StringBytes(e.G_a)
	x.Long(e.Key_fingerprint)
	x.Bytes(e.Protocol.Encode())
	return x.Buf
}

func (e TL_phone_setCallRating) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phone_setCallRating)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Rating)
	x.String(e.Comment)
	return x.Buf
}

func (e TL_phone_saveCallDebug) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_phone_saveCallDebug)
	x.Bytes(e.Peer.Encode())
	x.Bytes(e.Debug.Encode())
	return x.Buf
}

func (e TL_upload_getCdnFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_upload_getCdnFile)
	x.StringBytes(e.File_token)
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.Buf
}

func (e TL_upload_reuploadCdnFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_upload_reuploadCdnFile)
	x.StringBytes(e.File_token)
	x.StringBytes(e.Request_token)
	return x.Buf
}

func (e TL_help_getCdnConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_help_getCdnConfig)
	return x.Buf
}

func (e TL_messages_uploadMedia) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_uploadMedia)
	x.Bytes(e.Peer.Encode())
	x.Bytes(e.Media.Encode())
	return x.Buf
}

func (e TL_stickers_createStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_stickers_createStickerSet)
	x.Int(e.Flags)
	x.Bytes(e.User_id.Encode())
	x.String(e.Title)
	x.String(e.Short_name)
	x.Vector(e.Stickers)
	return x.Buf
}

func (e TL_langpack_getLangPack) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_langpack_getLangPack)
	x.String(e.Lang_code)
	return x.Buf
}

func (e TL_langpack_getStrings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_langpack_getStrings)
	x.String(e.Lang_code)
	x.VectorString(e.Keys)
	return x.Buf
}

func (e TL_langpack_getDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_langpack_getDifference)
	x.Int(e.From_version)
	return x.Buf
}

func (e TL_langpack_getLanguages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_langpack_getLanguages)
	return x.Buf
}

func (e TL_channels_editBanned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_editBanned)
	x.Bytes(e.Channel.Encode())
	x.Bytes(e.User_id.Encode())
	x.Bytes(e.Banned_rights.Encode())
	return x.Buf
}

func (e TL_channels_getAdminLog) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_getAdminLog)
	x.Int(e.Flags)
	x.Bytes(e.Channel.Encode())
	x.String(e.Q)
	x.Bytes(e.Events_filter.Encode())
	x.Vector(e.Admins)
	x.Long(e.Max_id)
	x.Long(e.Min_id)
	x.Int(e.Limit)
	return x.Buf
}

func (e TL_stickers_removeStickerFromSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_stickers_removeStickerFromSet)
	x.Bytes(e.Sticker.Encode())
	return x.Buf
}

func (e TL_stickers_changeStickerPosition) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_stickers_changeStickerPosition)
	x.Bytes(e.Sticker.Encode())
	x.Int(e.Position)
	return x.Buf
}

func (e TL_stickers_addStickerToSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_stickers_addStickerToSet)
	x.Bytes(e.Stickerset.Encode())
	x.Bytes(e.Sticker.Encode())
	return x.Buf
}

func (e TL_messages_sendScreenshotNotification) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_sendScreenshotNotification)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Reply_to_msg_id)
	x.Long(e.Random_id)
	return x.Buf
}

func (e TL_upload_getCdnFileHashes) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_upload_getCdnFileHashes)
	x.StringBytes(e.File_token)
	x.Int(e.Offset)
	return x.Buf
}

func (e TL_messages_getUnreadMentions) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getUnreadMentions)
	x.Bytes(e.Peer.Encode())
	x.Int(e.Offset_id)
	x.Int(e.Add_offset)
	x.Int(e.Limit)
	x.Int(e.Max_id)
	x.Int(e.Min_id)
	return x.Buf
}

func (e TL_messages_faveSticker) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_faveSticker)
	x.Bytes(e.Id.Encode())
	x.Bytes(e.Unfave.Encode())
	return x.Buf
}

func (e TL_channels_setStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_setStickers)
	x.Bytes(e.Channel.Encode())
	x.Bytes(e.Stickerset.Encode())
	return x.Buf
}

func (e TL_contacts_resetSaved) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_contacts_resetSaved)
	return x.Buf
}

func (e TL_messages_getFavedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_messages_getFavedStickers)
	x.Int(e.Hash)
	return x.Buf
}

func (e TL_channels_readMessageContents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc_channels_readMessageContents)
	x.Bytes(e.Channel.Encode())
	x.VectorInt(e.Id)
	return x.Buf
}

func (m *DecodeBuf) ObjectGenerated(constructor int32) (r TL) {
	fmt.Printf("!! ! ! ! ! ! ! !  ! ObjectGenerated: constructor: %d\n", constructor)
	switch constructor {
	case crc_boolFalse:
		r = TL_boolFalse{}

	case crc_boolTrue:
		r = TL_boolTrue{}

	case crc_error:
		r = TL_error{
			m.Int(),
			m.String(),
		}

	case crc_null:
		r = TL_null{}

	case crc_inputPeerEmpty:
		r = TL_inputPeerEmpty{}

	case crc_inputPeerSelf:
		r = TL_inputPeerSelf{}

	case crc_inputPeerChat:
		r = TL_inputPeerChat{
			m.Int(),
		}

	case crc_inputUserEmpty:
		r = TL_inputUserEmpty{}

	case crc_inputUserSelf:
		r = TL_inputUserSelf{}

	case crc_inputPhoneContact:
		r = TL_inputPhoneContact{
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_inputFile:
		r = TL_inputFile{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crc_inputMediaEmpty:
		r = TL_inputMediaEmpty{}

	case crc_inputMediaUploadedPhoto:
		flags := m.Flags()
		_ = flags
		r = TL_inputMediaUploadedPhoto{
			flags,
			m.Object(),
			m.String(),
			m.FlaggedVector(flags, 0),
			m.FlaggedInt(flags, 1),
		}

	case crc_inputMediaPhoto:
		flags := m.Flags()
		_ = flags
		r = TL_inputMediaPhoto{
			flags,
			m.Object(),
			m.String(),
			m.FlaggedInt(flags, 0),
		}

	case crc_inputMediaGeoPoint:
		r = TL_inputMediaGeoPoint{
			m.Object(),
		}

	case crc_inputMediaContact:
		r = TL_inputMediaContact{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_inputChatPhotoEmpty:
		r = TL_inputChatPhotoEmpty{}

	case crc_inputChatUploadedPhoto:
		r = TL_inputChatUploadedPhoto{
			m.Object(),
		}

	case crc_inputChatPhoto:
		r = TL_inputChatPhoto{
			m.Object(),
		}

	case crc_inputGeoPointEmpty:
		r = TL_inputGeoPointEmpty{}

	case crc_inputGeoPoint:
		r = TL_inputGeoPoint{
			m.Double(),
			m.Double(),
		}

	case crc_inputPhotoEmpty:
		r = TL_inputPhotoEmpty{}

	case crc_inputPhoto:
		r = TL_inputPhoto{
			m.Long(),
			m.Long(),
		}

	case crc_inputFileLocation:
		r = TL_inputFileLocation{
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case crc_inputAppEvent:
		r = TL_inputAppEvent{
			m.Double(),
			m.String(),
			m.Long(),
			m.String(),
		}

	case crc_peerUser:
		r = TL_peerUser{
			m.Int(),
		}

	case crc_peerChat:
		r = TL_peerChat{
			m.Int(),
		}

	case crc_storage_fileUnknown:
		r = TL_storage_fileUnknown{}

	case crc_storage_fileJpeg:
		r = TL_storage_fileJpeg{}

	case crc_storage_fileGif:
		r = TL_storage_fileGif{}

	case crc_storage_filePng:
		r = TL_storage_filePng{}

	case crc_storage_fileMp3:
		r = TL_storage_fileMp3{}

	case crc_storage_fileMov:
		r = TL_storage_fileMov{}

	case crc_storage_filePartial:
		r = TL_storage_filePartial{}

	case crc_storage_fileMp4:
		r = TL_storage_fileMp4{}

	case crc_storage_fileWebp:
		r = TL_storage_fileWebp{}

	case crc_fileLocationUnavailable:
		r = TL_fileLocationUnavailable{
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case crc_fileLocation:
		r = TL_fileLocation{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case crc_userEmpty:
		r = TL_userEmpty{
			m.Int(),
		}

	case crc_userProfilePhotoEmpty:
		r = TL_userProfilePhotoEmpty{}

	case crc_userProfilePhoto:
		r = TL_userProfilePhoto{
			m.Long(),
			m.Object(),
			m.Object(),
		}

	case crc_userStatusEmpty:
		r = TL_userStatusEmpty{}

	case crc_userStatusOnline:
		r = TL_userStatusOnline{
			m.Int(),
		}

	case crc_userStatusOffline:
		r = TL_userStatusOffline{
			m.Int(),
		}

	case crc_chatEmpty:
		r = TL_chatEmpty{
			m.Int(),
		}

	case crc_chat:
		flags := m.Flags()
		_ = flags
		r = TL_chat{
			flags,
			m.Int(),
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedObject(flags, 6),
		}

	case crc_chatForbidden:
		r = TL_chatForbidden{
			m.Int(),
			m.String(),
		}

	case crc_chatFull:
		r = TL_chatFull{
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Vector(),
		}

	case crc_chatParticipant:
		r = TL_chatParticipant{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_chatParticipantsForbidden:
		flags := m.Flags()
		_ = flags
		r = TL_chatParticipantsForbidden{
			flags,
			m.Int(),
			m.FlaggedObject(flags, 0),
		}

	case crc_chatParticipants:
		r = TL_chatParticipants{
			m.Int(),
			m.Vector(),
			m.Int(),
		}

	case crc_chatPhotoEmpty:
		r = TL_chatPhotoEmpty{}

	case crc_chatPhoto:
		r = TL_chatPhoto{
			m.Object(),
			m.Object(),
		}

	case crc_messageEmpty:
		r = TL_messageEmpty{
			m.Int(),
		}

	case crc_message:
		msg := TL_message{}

		flags := m.Flags()
		_ = flags

		msg.Flags = flags

		if (flags & (1 << 1)) != 0 {
			msg.Out = true
		}
		if (flags & (1 << 4)) != 0 {
			msg.Mentioned = true
		}
		if (flags & (1 << 5)) != 0 {
			msg.Media_unread = true
		}
		if (flags & (1 << 13)) != 0 {
			msg.Silent = true
		}
		if (flags & (1 << 14)) != 0 {
			msg.Post = true
		}

		msg.Id = m.Int()

		msg.From_id = m.FlaggedInt(flags, 8)

		msg.To_id = m.Object()

		msg.Fwd_from = m.FlaggedObject(flags, 2)

		msg.Via_bot_id = m.FlaggedInt(flags, 11)

		msg.Reply_to_msg_id = m.FlaggedInt(flags, 3)

		msg.Date = m.Int()
		msg.Message = m.String()

		msg.Media = m.FlaggedObject(flags, 9)

		msg.Reply_markup = m.FlaggedObject(flags, 6)

		msg.Entities = m.FlaggedVector(flags, 7)

		msg.Views = m.FlaggedInt(flags, 10)

		msg.Edit_date = m.FlaggedInt(flags, 15)

		msg.Post_author = m.FlaggedString(flags, 16)
		msg.GroupedId = m.FlaggedLong(flags, 17)

		msg.Likes = m.FlaggedInt(flags, 18)
		msg.Shares = m.FlaggedInt(flags, 19)
		msg.Comments = m.FlaggedInt(flags, 20)

		// r = TL_message{
		// 	flags,
		// 	m.Int(),
		// 	m.FlaggedInt(flags, 8),
		// 	m.Object(),
		// 	m.FlaggedObject(flags, 2),
		// 	m.FlaggedInt(flags, 11),
		// 	m.FlaggedInt(flags, 3),
		// 	m.Int(),
		// 	m.String(),
		// 	m.FlaggedObject(flags, 9),
		// 	m.FlaggedObject(flags, 6),
		// 	m.FlaggedVector(flags, 7),
		// 	m.FlaggedInt(flags, 10),
		// 	m.FlaggedInt(flags, 15),
		// 	m.FlaggedString(flags, 16),
		// }
	case crc_messageService:
		flags := m.Flags()
		_ = flags
		r = TL_messageService{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 8),
			m.Object(),
			m.FlaggedInt(flags, 3),
			m.Int(),
			m.Object(),
		}

	case crc_messageMediaEmpty:
		r = TL_messageMediaEmpty{}

	case crc_messageMediaPhoto:
		flags := m.Flags()
		_ = flags
		r = TL_messageMediaPhoto{
			flags,
			m.FlaggedObject(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedInt(flags, 2),
		}

	case crc_messageMediaGeo:
		r = TL_messageMediaGeo{
			m.Object(),
		}

	case crc_messageMediaContact:
		r = TL_messageMediaContact{
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
		}

	case crc_messageMediaUnsupported:
		r = TL_messageMediaUnsupported{}

	case crc_messageActionEmpty:
		r = TL_messageActionEmpty{}

	case crc_messageActionChatCreate:
		r = TL_messageActionChatCreate{
			m.String(),
			m.VectorInt(),
		}

	case crc_messageActionChatEditTitle:
		r = TL_messageActionChatEditTitle{
			m.String(),
		}

	case crc_messageActionChatEditPhoto:
		r = TL_messageActionChatEditPhoto{
			m.Object(),
		}

	case crc_messageActionChatDeletePhoto:
		r = TL_messageActionChatDeletePhoto{}

	case crc_messageActionChatAddUser:
		r = TL_messageActionChatAddUser{
			m.VectorInt(),
		}

	case crc_messageActionChatDeleteUser:
		r = TL_messageActionChatDeleteUser{
			m.Int(),
		}

	case crc_dialog:
		flags := m.Flags()
		_ = flags
		r = TL_dialog{
			flags,
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.FlaggedObject(flags, 1),
		}

	case crc_photoEmpty:
		r = TL_photoEmpty{
			m.Long(),
		}

	case crc_photo:
		flags := m.Flags()
		_ = flags
		r = TL_photo{
			flags,
			m.Long(),
			m.Long(),
			m.Int(),
			m.Vector(),
		}

	case crc_photoSizeEmpty:
		r = TL_photoSizeEmpty{
			m.String(),
		}

	case crc_photoSize:
		r = TL_photoSize{
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_photoCachedSize:
		r = TL_photoCachedSize{
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_geoPointEmpty:
		r = TL_geoPointEmpty{}

	case crc_geoPoint:
		r = TL_geoPoint{
			m.Double(),
			m.Double(),
		}

	case crc_auth_checkedPhone:
		r = TL_auth_checkedPhone{
			m.Object(),
		}

	case crc_auth_sentCode:
		flags := m.Flags()

		aS := TL_auth_sentCode{}
		if (flags & (1 << 0)) != 0 {
			aS.PhoneRegistered = true
		}

		aS.Type = m.Object()
		aS.PhoneCodeHash = m.String()

		if (flags & (1 << 1)) != 0 {
			aS.NextType = m.Object()
		}
		if (flags & (1 << 2)) != 0 {
			aS.Timeout = m.Int()
		}
		if (flags & (1 << 3)) != 0 {
			aS.TermsOfService = m.Object()
		}

		r = aS
		// _ = flags
		// r = TL_auth_sentCode{
		// 	flags,
		// 	m.Object(),
		// 	m.String(),
		// 	m.FlaggedObject(flags, 1),
		// 	m.FlaggedInt(flags, 2),
		// }

	case crc_auth_authorization:
		flags := m.Flags()
		_ = flags
		r = TL_auth_authorization{
			flags,
			m.FlaggedInt(flags, 0),
			m.Object(),
		}

	case crc_auth_exportedAuthorization:
		r = TL_auth_exportedAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case crc_inputNotifyPeer:
		r = TL_inputNotifyPeer{
			m.Object(),
		}

	case crc_inputNotifyUsers:
		r = TL_inputNotifyUsers{}

	case crc_inputNotifyChats:
		r = TL_inputNotifyChats{}

	case crc_inputNotifyAll:
		r = TL_inputNotifyAll{}

	case crc_inputPeerNotifySettings:
		flags := m.Flags()
		_ = flags
		r = TL_inputPeerNotifySettings{
			flags,
			m.Int(),
			m.String(),
		}

	case crc_peerNotifyEventsEmpty:
		r = TL_peerNotifyEventsEmpty{}

	case crc_peerNotifyEventsAll:
		r = TL_peerNotifyEventsAll{}

	case crc_peerNotifySettingsEmpty:
		r = TL_peerNotifySettingsEmpty{}

	case crc_peerNotifySettings:
		flags := m.Flags()
		_ = flags
		r = TL_peerNotifySettings{
			flags,
			m.Int(),
			m.String(),
		}

	case crc_wallPaper:
		r = TL_wallPaper{
			m.Int(),
			m.String(),
			m.Vector(),
			m.Int(),
		}

	case crc_userFull:
		flags := m.Flags()
		_ = flags
		r = TL_userFull{
			flags,
			m.Object(),
			m.FlaggedString(flags, 1),
			m.Object(),
			m.FlaggedObject(flags, 2),
			m.Object(),
			m.FlaggedObject(flags, 3),
			m.Int(),
		}

	case crc_contact:
		r = TL_contact{
			m.Int(),
			m.Object(),
		}

	case crc_importedContact:
		r = TL_importedContact{
			m.Int(),
			m.Long(),
		}

	case crc_contactBlocked:
		r = TL_contactBlocked{
			m.Int(),
			m.Int(),
		}

	case crc_contactStatus:
		r = TL_contactStatus{
			m.Int(),
			m.Object(),
		}

	case crc_contacts_link:
		r = TL_contacts_link{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_contacts_contacts:
		r = TL_contacts_contacts{
			m.Vector(),
			m.Int(),
			m.Vector(),
		}

	case crc_contacts_contactsNotModified:
		r = TL_contacts_contactsNotModified{}

	case crc_contacts_importedContacts:
		r = TL_contacts_importedContacts{
			m.Vector(),
			m.Vector(),
			m.VectorLong(),
			m.Vector(),
		}

	case crc_contacts_blocked:
		r = TL_contacts_blocked{
			m.Vector(),
			m.Vector(),
		}

	case crc_contacts_blockedSlice:
		r = TL_contacts_blockedSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crc_contacts_found:
		r = TL_contacts_found{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_dialogs:
		r = TL_messages_dialogs{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_dialogsSlice:
		r = TL_messages_dialogsSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_messages:
		r = TL_messages_messages{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_messagesSlice:
		r = TL_messages_messagesSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_chats:
		r = TL_messages_chats{
			m.Vector(),
		}

	case crc_messages_chatFull:
		r = TL_messages_chatFull{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_affectedHistory:
		r = TL_messages_affectedHistory{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputMessagesFilterEmpty:
		r = TL_inputMessagesFilterEmpty{}

	case crc_inputMessagesFilterPhotos:
		r = TL_inputMessagesFilterPhotos{}

	case crc_inputMessagesFilterVideo:
		r = TL_inputMessagesFilterVideo{}

	case crc_inputMessagesFilterPhotoVideo:
		r = TL_inputMessagesFilterPhotoVideo{}

	case crc_updateNewMessage:
		r = TL_updateNewMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_updateMessageID:
		r = TL_updateMessageID{
			m.Int(),
			m.Long(),
		}

	case crc_updateDeleteMessages:
		r = TL_updateDeleteMessages{
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crc_updateUserTyping:
		r = TL_updateUserTyping{
			m.Int(),
			m.Object(),
		}

	case crc_updateChatUserTyping:
		r = TL_updateChatUserTyping{
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_updateChatParticipants:
		r = TL_updateChatParticipants{
			m.Object(),
		}

	case crc_updateUserStatus:
		r = TL_updateUserStatus{
			m.Int(),
			m.Object(),
		}

	case crc_updateUserName:
		r = TL_updateUserName{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_updateUserPhoto:
		r = TL_updateUserPhoto{
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crc_updateContactRegistered:
		r = TL_updateContactRegistered{
			m.Int(),
			m.Int(),
		}

	case crc_updateContactLink:
		r = TL_updateContactLink{
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crc_updates_state:
		r = TL_updates_state{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updates_differenceEmpty:
		r = TL_updates_differenceEmpty{
			m.Int(),
			m.Int(),
		}

	case crc_updates_difference:
		r = TL_updates_difference{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crc_updates_differenceSlice:
		r = TL_updates_differenceSlice{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crc_updatesTooLong:
		r = TL_updatesTooLong{}

	case crc_updateShortMessage:
		flags := m.Flags()
		_ = flags
		r = TL_updateShortMessage{
			flags,
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedObject(flags, 2),
			m.FlaggedInt(flags, 11),
			m.FlaggedInt(flags, 3),
			m.FlaggedVector(flags, 7),
		}

	case crc_updateShortChatMessage:
		flags := m.Flags()
		_ = flags
		r = TL_updateShortChatMessage{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedObject(flags, 2),
			m.FlaggedInt(flags, 11),
			m.FlaggedInt(flags, 3),
			m.FlaggedVector(flags, 7),
		}

	case crc_updateShort:
		r = TL_updateShort{
			m.Object(),
			m.Int(),
		}

	case crc_updatesCombined:
		r = TL_updatesCombined{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updates:
		r = TL_updates{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case crc_photos_photo:
		r = TL_photos_photo{
			m.Object(),
			m.Vector(),
		}

	case crc_upload_file:
		r = TL_upload_file{
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_dcOption:
		dc := TL_dcOption{}
		flags := m.Flags()
		_ = flags
		if (flags & (1 << 0)) != 0 {
			dc.Ipv6 = true
		}
		if (flags & (1 << 1)) != 0 {
			dc.MediaOnly = true
		}
		if (flags & (1 << 2)) != 0 {
			dc.TCPoOnly = true
		}
		if (flags & (1 << 3)) != 0 {
			dc.Cdn = true
		}
		if (flags & (1 << 4)) != 0 {
			dc.Static = true
		}
		dc.Flags = flags
		dc.Id = m.Int()
		dc.IpAddress = m.String()
		dc.Port = m.Int()

		if (flags & (1 << 10)) != 0 {
			dc.Secret = m.StringBytes()
		}

		r = dc

	case crc_config:
		c := TL_config{}
		flags := m.Flags()
		_ = flags

		c.Flags = flags
		c.Date = m.Int()
		c.Expires = m.Int()
		c.TestMode = m.Object()
		c.ThisDC = m.Int()
		c.DcOptions = m.Vector()
		c.ChatSizeMax = m.Int()
		c.MegagroupSizeMax = m.Int()
		c.ForwardedCountMax = m.Int()
		c.OnlineUpdatePeriodMs = m.Int()
		c.OfflineBlurTimeoutMs = m.Int()
		c.OfflineIdleTimeoutMs = m.Int()
		c.OnlineCloudTimeoutMs = m.Int()
		c.NotifyCloudDelayMs = m.Int()
		c.NotifyDefaultDelayMs = m.Int()
		c.ChatBigSize = m.Int()
		c.PushChatPeriodMs = m.Int()
		c.PushChatLimit = m.Int()
		c.SavedGifsLimit = m.Int()
		c.EditTimeLimit = m.Int()
		c.RatingEDecay = m.Int()
		c.StickersRecentLimit = m.Int()
		c.StickersFavedLimit = m.Int()
		c.ChannelsReadMediaPeriod = m.Int()
		// m.FlaggedInt(flags, 0),
		if (flags & (1 << 0)) != 0 {
			c.TmpSessions = m.Int()
		}
		c.PinnedDialogsCountMax = m.Int()
		c.CallReceiveTimeoutMs = m.Int()
		c.CallRingTimeoutMs = m.Int()
		c.CallConnectTimeoutMs = m.Int()
		c.CallPacketTimeoutMs = m.Int()
		c.MeUrlPrefix = m.String()
		if (flags & (1 << 2)) != 0 {
			c.SuggestedLangCode = m.String()
		}
		if (flags & (1 << 2)) != 0 {
			c.LangPackVersion = m.Int()
		}
		c.Magic = m.Int()
		c.CountDisableFeature = m.Int()

		r = c

	case crc_nearestDc:
		r = TL_nearestDc{
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_help_appUpdate:
		r = TL_help_appUpdate{
			m.Int(),
			m.Object(),
			m.String(),
			m.String(),
		}

	case crc_help_noAppUpdate:
		r = TL_help_noAppUpdate{}

	case crc_help_inviteText:
		r = TL_help_inviteText{
			m.String(),
		}

	case crc_inputPeerNotifyEventsEmpty:
		r = TL_inputPeerNotifyEventsEmpty{}

	case crc_inputPeerNotifyEventsAll:
		r = TL_inputPeerNotifyEventsAll{}

	case crc_photos_photos:
		r = TL_photos_photos{
			m.Vector(),
			m.Vector(),
		}

	case crc_photos_photosSlice:
		r = TL_photos_photosSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crc_wallPaperSolid:
		r = TL_wallPaperSolid{
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_updateNewEncryptedMessage:
		r = TL_updateNewEncryptedMessage{
			m.Object(),
			m.Int(),
		}

	case crc_updateEncryptedChatTyping:
		r = TL_updateEncryptedChatTyping{
			m.Int(),
		}

	case crc_updateEncryption:
		r = TL_updateEncryption{
			m.Object(),
			m.Int(),
		}

	case crc_updateEncryptedMessagesRead:
		r = TL_updateEncryptedMessagesRead{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_encryptedChatEmpty:
		r = TL_encryptedChatEmpty{
			m.Int(),
		}

	case crc_encryptedChatWaiting:
		r = TL_encryptedChatWaiting{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_encryptedChatRequested:
		r = TL_encryptedChatRequested{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_encryptedChat:
		r = TL_encryptedChat{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Long(),
		}

	case crc_encryptedChatDiscarded:
		r = TL_encryptedChatDiscarded{
			m.Int(),
		}

	case crc_inputEncryptedChat:
		r = TL_inputEncryptedChat{
			m.Int(),
			m.Long(),
		}

	case crc_encryptedFileEmpty:
		r = TL_encryptedFileEmpty{}

	case crc_encryptedFile:
		r = TL_encryptedFile{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputEncryptedFileEmpty:
		r = TL_inputEncryptedFileEmpty{}

	case crc_inputEncryptedFileUploaded:
		r = TL_inputEncryptedFileUploaded{
			m.Long(),
			m.Int(),
			m.String(),
			m.Int(),
		}

	case crc_inputEncryptedFile:
		r = TL_inputEncryptedFile{
			m.Long(),
			m.Long(),
		}

	case crc_inputEncryptedFileLocation:
		r = TL_inputEncryptedFileLocation{
			m.Long(),
			m.Long(),
		}

	case crc_encryptedMessage:
		r = TL_encryptedMessage{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_encryptedMessageService:
		r = TL_encryptedMessageService{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messages_dhConfigNotModified:
		r = TL_messages_dhConfigNotModified{
			m.StringBytes(),
		}

	case crc_messages_dhConfig:
		r = TL_messages_dhConfig{
			m.Int(),
			m.StringBytes(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messages_sentEncryptedMessage:
		r = TL_messages_sentEncryptedMessage{
			m.Int(),
		}

	case crc_messages_sentEncryptedFile:
		r = TL_messages_sentEncryptedFile{
			m.Int(),
			m.Object(),
		}

	case crc_inputFileBig:
		r = TL_inputFileBig{
			m.Long(),
			m.Int(),
			m.String(),
		}

	case crc_inputEncryptedFileBigUploaded:
		r = TL_inputEncryptedFileBigUploaded{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case crc_storage_filePdf:
		r = TL_storage_filePdf{}

	case crc_inputMessagesFilterDocument:
		r = TL_inputMessagesFilterDocument{}

	case crc_inputMessagesFilterPhotoVideoDocuments:
		r = TL_inputMessagesFilterPhotoVideoDocuments{}

	case crc_updateChatParticipantAdd:
		r = TL_updateChatParticipantAdd{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChatParticipantDelete:
		r = TL_updateChatParticipantDelete{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateDcOptions:
		r = TL_updateDcOptions{
			m.Vector(),
		}

	case crc_inputMediaUploadedDocument:
		flags := m.Flags()
		_ = flags
		r = TL_inputMediaUploadedDocument{
			flags,
			m.Object(),
			m.FlaggedObject(flags, 2),
			m.String(),
			m.Vector(),
			m.String(),
			m.FlaggedVector(flags, 0),
			m.FlaggedInt(flags, 1),
		}

	case crc_inputMediaDocument:
		flags := m.Flags()
		_ = flags
		r = TL_inputMediaDocument{
			flags,
			m.Object(),
			m.String(),
			m.FlaggedInt(flags, 0),
		}

	case crc_messageMediaDocument:
		flags := m.Flags()
		_ = flags
		r = TL_messageMediaDocument{
			flags,
			m.FlaggedObject(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedInt(flags, 2),
		}

	case crc_inputDocumentEmpty:
		r = TL_inputDocumentEmpty{}

	case crc_inputDocument:
		r = TL_inputDocument{
			m.Long(),
			m.Long(),
		}

	case crc_inputDocumentFileLocation:
		r = TL_inputDocumentFileLocation{
			m.Long(),
			m.Long(),
			m.Int(),
		}

	case crc_documentEmpty:
		r = TL_documentEmpty{
			m.Long(),
		}

	case crc_document:
		r = TL_document{
			m.Long(),
			m.Long(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Vector(),
		}

	case crc_help_support:
		r = TL_help_support{
			m.String(),
			m.Object(),
		}

	case crc_notifyAll:
		r = TL_notifyAll{}

	case crc_notifyChats:
		r = TL_notifyChats{}

	case crc_notifyPeer:
		r = TL_notifyPeer{
			m.Object(),
		}

	case crc_notifyUsers:
		r = TL_notifyUsers{}

	case crc_updateUserBlocked:
		r = TL_updateUserBlocked{
			m.Int(),
			m.Object(),
		}

	case crc_updateNotifySettings:
		r = TL_updateNotifySettings{
			m.Object(),
			m.Object(),
		}

	case crc_sendMessageTypingAction:
		r = TL_sendMessageTypingAction{}

	case crc_sendMessageCancelAction:
		r = TL_sendMessageCancelAction{}

	case crc_sendMessageRecordVideoAction:
		r = TL_sendMessageRecordVideoAction{}

	case crc_sendMessageUploadVideoAction:
		r = TL_sendMessageUploadVideoAction{
			m.Int(),
		}

	case crc_sendMessageRecordAudioAction:
		r = TL_sendMessageRecordAudioAction{}

	case crc_sendMessageUploadAudioAction:
		r = TL_sendMessageUploadAudioAction{
			m.Int(),
		}

	case crc_sendMessageUploadPhotoAction:
		r = TL_sendMessageUploadPhotoAction{
			m.Int(),
		}

	case crc_sendMessageUploadDocumentAction:
		r = TL_sendMessageUploadDocumentAction{
			m.Int(),
		}

	case crc_sendMessageGeoLocationAction:
		r = TL_sendMessageGeoLocationAction{}

	case crc_sendMessageChooseContactAction:
		r = TL_sendMessageChooseContactAction{}

	case crc_updateServiceNotification:
		flags := m.Flags()
		_ = flags
		r = TL_updateServiceNotification{
			flags,
			m.FlaggedInt(flags, 1),
			m.String(),
			m.String(),
			m.Object(),
			m.Vector(),
		}

	case crc_userStatusRecently:
		r = TL_userStatusRecently{}

	case crc_userStatusLastWeek:
		r = TL_userStatusLastWeek{}

	case crc_userStatusLastMonth:
		r = TL_userStatusLastMonth{}

	case crc_updatePrivacy:
		r = TL_updatePrivacy{
			m.Object(),
			m.Vector(),
		}

	case crc_inputPrivacyKeyStatusTimestamp:
		r = TL_inputPrivacyKeyStatusTimestamp{}

	case crc_privacyKeyStatusTimestamp:
		r = TL_privacyKeyStatusTimestamp{}

	case crc_inputPrivacyValueAllowContacts:
		r = TL_inputPrivacyValueAllowContacts{}

	case crc_inputPrivacyValueAllowAll:
		r = TL_inputPrivacyValueAllowAll{}

	case crc_inputPrivacyValueAllowUsers:
		r = TL_inputPrivacyValueAllowUsers{
			m.Vector(),
		}

	case crc_inputPrivacyValueDisallowContacts:
		r = TL_inputPrivacyValueDisallowContacts{}

	case crc_inputPrivacyValueDisallowAll:
		r = TL_inputPrivacyValueDisallowAll{}

	case crc_inputPrivacyValueDisallowUsers:
		r = TL_inputPrivacyValueDisallowUsers{
			m.Vector(),
		}

	case crc_privacyValueAllowContacts:
		r = TL_privacyValueAllowContacts{}

	case crc_privacyValueAllowAll:
		r = TL_privacyValueAllowAll{}

	case crc_privacyValueAllowUsers:
		r = TL_privacyValueAllowUsers{
			m.VectorInt(),
		}

	case crc_privacyValueDisallowContacts:
		r = TL_privacyValueDisallowContacts{}

	case crc_privacyValueDisallowAll:
		r = TL_privacyValueDisallowAll{}

	case crc_privacyValueDisallowUsers:
		r = TL_privacyValueDisallowUsers{
			m.VectorInt(),
		}

	case crc_account_privacyRules:
		r = TL_account_privacyRules{
			m.Vector(),
			m.Vector(),
		}

	case crc_accountDaysTTL:
		r = TL_accountDaysTTL{
			m.Int(),
		}

	case crc_updateUserPhone:
		r = TL_updateUserPhone{
			m.Int(),
			m.String(),
		}

	case crc_disabledFeature:
		r = TL_disabledFeature{
			m.String(),
			m.String(),
		}

	case crc_documentAttributeImageSize:
		r = TL_documentAttributeImageSize{
			m.Int(),
			m.Int(),
		}

	case crc_documentAttributeAnimated:
		r = TL_documentAttributeAnimated{}

	case crc_documentAttributeSticker:
		flags := m.Flags()
		_ = flags
		r = TL_documentAttributeSticker{
			flags,
			m.String(),
			m.Object(),
			m.FlaggedObject(flags, 0),
		}

	case crc_documentAttributeVideo:
		flags := m.Flags()
		_ = flags
		r = TL_documentAttributeVideo{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_documentAttributeAudio:
		flags := m.Flags()
		_ = flags
		r = TL_documentAttributeAudio{
			flags,
			m.Int(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedStringBytes(flags, 2),
		}

	case crc_documentAttributeFilename:
		r = TL_documentAttributeFilename{
			m.String(),
		}

	case crc_messages_stickersNotModified:
		r = TL_messages_stickersNotModified{}

	case crc_messages_stickers:
		r = TL_messages_stickers{
			m.String(),
			m.Vector(),
		}

	case crc_stickerPack:
		r = TL_stickerPack{
			m.String(),
			m.VectorLong(),
		}

	case crc_messages_allStickersNotModified:
		r = TL_messages_allStickersNotModified{}

	case crc_messages_allStickers:
		r = TL_messages_allStickers{
			m.Int(),
			m.Vector(),
		}

	case crc_account_noPassword:
		r = TL_account_noPassword{
			m.StringBytes(),
			m.String(),
		}

	case crc_account_password:
		r = TL_account_password{
			m.StringBytes(),
			m.StringBytes(),
			m.String(),
			m.Object(),
			m.String(),
		}

	case crc_updateReadHistoryInbox:
		r = TL_updateReadHistoryInbox{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateReadHistoryOutbox:
		r = TL_updateReadHistoryOutbox{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_affectedMessages:
		r = TL_messages_affectedMessages{
			m.Int(),
			m.Int(),
		}

	case crc_contactLinkUnknown:
		r = TL_contactLinkUnknown{}

	case crc_contactLinkNone:
		r = TL_contactLinkNone{}

	case crc_contactLinkHasPhone:
		r = TL_contactLinkHasPhone{}

	case crc_contactLinkContact:
		r = TL_contactLinkContact{}

	case crc_updateWebPage:
		r = TL_updateWebPage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_webPageEmpty:
		r = TL_webPageEmpty{
			m.Long(),
		}

	case crc_webPagePending:
		r = TL_webPagePending{
			m.Long(),
			m.Int(),
		}

	case crc_webPage:
		flags := m.Flags()
		_ = flags
		r = TL_webPage{
			flags,
			m.Long(),
			m.String(),
			m.String(),
			m.Int(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedObject(flags, 4),
			m.FlaggedString(flags, 5),
			m.FlaggedString(flags, 5),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 7),
			m.FlaggedString(flags, 8),
			m.FlaggedObject(flags, 9),
			m.FlaggedObject(flags, 10),
		}

	case crc_messageMediaWebPage:
		r = TL_messageMediaWebPage{
			m.Object(),
		}

	case crc_authorization:
		r = TL_authorization{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_account_authorizations:
		r = TL_account_authorizations{
			m.Vector(),
		}

	case crc_account_passwordSettings:
		r = TL_account_passwordSettings{
			m.String(),
		}

	case crc_account_passwordInputSettings:
		flags := m.Flags()
		_ = flags
		r = TL_account_passwordInputSettings{
			flags,
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
		}

	case crc_auth_passwordRecovery:
		r = TL_auth_passwordRecovery{
			m.String(),
		}

	case crc_inputMediaVenue:
		r = TL_inputMediaVenue{
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_messageMediaVenue:
		r = TL_messageMediaVenue{
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_receivedNotifyMessage:
		r = TL_receivedNotifyMessage{
			m.Int(),
			m.Int(),
		}

	case crc_chatInviteEmpty:
		r = TL_chatInviteEmpty{}

	case crc_chatInviteExported:
		r = TL_chatInviteExported{
			m.String(),
		}

	case crc_chatInviteAlready:
		r = TL_chatInviteAlready{
			m.Object(),
		}

	case crc_chatInvite:
		flags := m.Flags()
		_ = flags
		r = TL_chatInvite{
			flags,
			m.String(),
			m.Object(),
			m.Int(),
			m.FlaggedVector(flags, 4),
		}

	case crc_messageActionChatJoinedByLink:
		r = TL_messageActionChatJoinedByLink{
			m.Int(),
		}

	case crc_updateReadMessagesContents:
		r = TL_updateReadMessagesContents{
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crc_inputStickerSetEmpty:
		r = TL_inputStickerSetEmpty{}

	case crc_inputStickerSetID:
		r = TL_inputStickerSetID{
			m.Long(),
			m.Long(),
		}

	case crc_inputStickerSetShortName:
		r = TL_inputStickerSetShortName{
			m.String(),
		}

	case crc_stickerSet:
		flags := m.Flags()
		_ = flags
		r = TL_stickerSet{
			flags,
			m.Long(),
			m.Long(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_stickerSet:
		r = TL_messages_stickerSet{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case crc_user:
		flags := m.Flags()
		_ = flags
		r = TL_user{
			flags,
			m.Int(),
			m.FlaggedLong(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedString(flags, 4),
			m.FlaggedObject(flags, 5),
			m.FlaggedObject(flags, 6),
			m.FlaggedInt(flags, 14),
			m.FlaggedString(flags, 18),
			m.FlaggedString(flags, 19),
			m.FlaggedString(flags, 22),
		}

	case crc_botCommand:
		r = TL_botCommand{
			m.String(),
			m.String(),
		}

	case crc_botInfo:
		r = TL_botInfo{
			m.Int(),
			m.String(),
			m.Vector(),
		}

	case crc_keyboardButton:
		r = TL_keyboardButton{
			m.String(),
		}

	case crc_keyboardButtonRow:
		r = TL_keyboardButtonRow{
			m.Vector(),
		}

	case crc_replyKeyboardHide:
		flags := m.Flags()
		_ = flags
		r = TL_replyKeyboardHide{
			flags,
		}

	case crc_replyKeyboardForceReply:
		flags := m.Flags()
		_ = flags
		r = TL_replyKeyboardForceReply{
			flags,
		}

	case crc_replyKeyboardMarkup:
		flags := m.Flags()
		_ = flags
		r = TL_replyKeyboardMarkup{
			flags,
			m.Vector(),
		}

	case crc_inputMessagesFilterUrl:
		r = TL_inputMessagesFilterUrl{}

	case crc_inputPeerUser:
		r = TL_inputPeerUser{
			m.Int(),
			m.Long(),
		}

	case crc_inputUser:
		r = TL_inputUser{
			m.Int(),
			m.Long(),
		}

	case crc_messageEntityUnknown:
		r = TL_messageEntityUnknown{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityMention:
		r = TL_messageEntityMention{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityHashtag:
		r = TL_messageEntityHashtag{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityBotCommand:
		r = TL_messageEntityBotCommand{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityUrl:
		r = TL_messageEntityUrl{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityEmail:
		r = TL_messageEntityEmail{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityBold:
		r = TL_messageEntityBold{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityItalic:
		r = TL_messageEntityItalic{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityCode:
		r = TL_messageEntityCode{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityPre:
		r = TL_messageEntityPre{
			m.Int(),
			m.Int(),
			m.String(),
		}

	case crc_messageEntityTextUrl:
		r = TL_messageEntityTextUrl{
			m.Int(),
			m.Int(),
			m.String(),
		}

	case crc_updateShortSentMessage:
		flags := m.Flags()
		_ = flags
		r = TL_updateShortSentMessage{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedObject(flags, 9),
			m.FlaggedVector(flags, 7),
		}

	case crc_inputPeerChannel:
		r = TL_inputPeerChannel{
			m.Int(),
			m.Long(),
		}

	case crc_peerChannel:
		r = TL_peerChannel{
			m.Int(),
		}

	case crc_channel:
		flags := m.Flags()
		_ = flags
		r = TL_channel{
			flags,
			m.Int(),
			m.FlaggedLong(flags, 13),
			m.String(),
			m.FlaggedString(flags, 6),
			m.Object(),
			m.Int(),
			m.Int(),
			m.FlaggedString(flags, 9),
			m.FlaggedObject(flags, 14),
			m.FlaggedObject(flags, 15),
		}

	case crc_channelForbidden:
		flags := m.Flags()
		_ = flags
		r = TL_channelForbidden{
			flags,
			m.Int(),
			m.Long(),
			m.String(),
			m.FlaggedInt(flags, 16),
		}

	case crc_channelFull:
		flags := m.Flags()
		_ = flags
		r = TL_channelFull{
			flags,
			m.Int(),
			m.String(),
			m.FlaggedInt(flags, 0),
			m.FlaggedInt(flags, 1),
			m.FlaggedInt(flags, 2),
			m.FlaggedInt(flags, 2),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Vector(),
			m.FlaggedInt(flags, 4),
			m.FlaggedInt(flags, 4),
			m.FlaggedInt(flags, 5),
			m.FlaggedObject(flags, 8),
		}

	case crc_messageActionChannelCreate:
		r = TL_messageActionChannelCreate{
			m.String(),
		}

	case crc_messages_channelMessages:
		flags := m.Flags()
		_ = flags
		r = TL_messages_channelMessages{
			flags,
			m.Int(),
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_updateChannelTooLong:
		flags := m.Flags()
		_ = flags
		r = TL_updateChannelTooLong{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 0),
		}

	case crc_updateChannel:
		r = TL_updateChannel{
			m.Int(),
		}

	case crc_updateNewChannelMessage:
		r = TL_updateNewChannelMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_updateReadChannelInbox:
		r = TL_updateReadChannelInbox{
			m.Int(),
			m.Int(),
		}

	case crc_updateDeleteChannelMessages:
		r = TL_updateDeleteChannelMessages{
			m.Int(),
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChannelMessageViews:
		r = TL_updateChannelMessageViews{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputChannelEmpty:
		r = TL_inputChannelEmpty{}

	case crc_inputChannel:
		r = TL_inputChannel{
			m.Int(),
			m.Long(),
		}

	case crc_contacts_resolvedPeer:
		r = TL_contacts_resolvedPeer{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messageRange:
		r = TL_messageRange{
			m.Int(),
			m.Int(),
		}

	case crc_updates_channelDifferenceEmpty:
		flags := m.Flags()
		_ = flags
		r = TL_updates_channelDifferenceEmpty{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 1),
		}

	case crc_updates_channelDifferenceTooLong:
		flags := m.Flags()
		_ = flags
		r = TL_updates_channelDifferenceTooLong{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 1),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_updates_channelDifference:
		flags := m.Flags()
		_ = flags
		r = TL_updates_channelDifference{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 1),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_channelMessagesFilterEmpty:
		r = TL_channelMessagesFilterEmpty{}

	case crc_channelMessagesFilter:
		flags := m.Flags()
		_ = flags
		r = TL_channelMessagesFilter{
			flags,
			m.Vector(),
		}

	case crc_channelParticipant:
		r = TL_channelParticipant{
			m.Int(),
			m.Int(),
		}

	case crc_channelParticipantSelf:
		r = TL_channelParticipantSelf{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_channelParticipantCreator:
		r = TL_channelParticipantCreator{
			m.Int(),
		}

	case crc_channelParticipantsRecent:
		r = TL_channelParticipantsRecent{}

	case crc_channelParticipantsAdmins:
		r = TL_channelParticipantsAdmins{}

	case crc_channelParticipantsKicked:
		r = TL_channelParticipantsKicked{
			m.String(),
		}

	case crc_channels_channelParticipants:
		r = TL_channels_channelParticipants{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crc_channels_channelParticipant:
		r = TL_channels_channelParticipant{
			m.Object(),
			m.Vector(),
		}

	case crc_true:
		r = TL_true{}

	case crc_chatParticipantCreator:
		r = TL_chatParticipantCreator{
			m.Int(),
		}

	case crc_chatParticipantAdmin:
		r = TL_chatParticipantAdmin{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChatAdmins:
		r = TL_updateChatAdmins{
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_updateChatParticipantAdmin:
		r = TL_updateChatParticipantAdmin{
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_messageActionChatMigrateTo:
		r = TL_messageActionChatMigrateTo{
			m.Int(),
		}

	case crc_messageActionChannelMigrateFrom:
		r = TL_messageActionChannelMigrateFrom{
			m.String(),
			m.Int(),
		}

	case crc_channelParticipantsBots:
		r = TL_channelParticipantsBots{}

	case crc_inputReportReasonSpam:
		r = TL_inputReportReasonSpam{}

	case crc_inputReportReasonViolence:
		r = TL_inputReportReasonViolence{}

	case crc_inputReportReasonPornography:
		r = TL_inputReportReasonPornography{}

	case crc_inputReportReasonOther:
		r = TL_inputReportReasonOther{
			m.String(),
		}

	case crc_updateNewStickerSet:
		r = TL_updateNewStickerSet{
			m.Object(),
		}

	case crc_updateStickerSetsOrder:
		flags := m.Flags()
		_ = flags
		r = TL_updateStickerSetsOrder{
			flags,
			m.VectorLong(),
		}

	case crc_updateStickerSets:
		r = TL_updateStickerSets{}

	case crc_help_termsOfService:
		r = TL_help_termsOfService{
			m.String(),
		}

	case crc_foundGif:
		r = TL_foundGif{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_inputMediaGifExternal:
		r = TL_inputMediaGifExternal{
			m.String(),
			m.String(),
		}

	case crc_messages_foundGifs:
		r = TL_messages_foundGifs{
			m.Int(),
			m.Vector(),
		}

	case crc_inputMessagesFilterGif:
		r = TL_inputMessagesFilterGif{}

	case crc_updateSavedGifs:
		r = TL_updateSavedGifs{}

	case crc_updateBotInlineQuery:
		flags := m.Flags()
		_ = flags
		r = TL_updateBotInlineQuery{
			flags,
			m.Long(),
			m.Int(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.String(),
		}

	case crc_foundGifCached:
		r = TL_foundGifCached{
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_savedGifsNotModified:
		r = TL_messages_savedGifsNotModified{}

	case crc_messages_savedGifs:
		r = TL_messages_savedGifs{
			m.Int(),
			m.Vector(),
		}

	case crc_inputBotInlineMessageMediaAuto:
		flags := m.Flags()
		_ = flags
		r = TL_inputBotInlineMessageMediaAuto{
			flags,
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case crc_inputBotInlineMessageText:
		flags := m.Flags()
		_ = flags
		r = TL_inputBotInlineMessageText{
			flags,
			m.String(),
			m.FlaggedVector(flags, 1),
			m.FlaggedObject(flags, 2),
		}

	case crc_inputBotInlineResult:
		flags := m.Flags()
		_ = flags
		r = TL_inputBotInlineResult{
			flags,
			m.String(),
			m.String(),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedString(flags, 4),
			m.FlaggedString(flags, 5),
			m.FlaggedString(flags, 5),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 7),
			m.Object(),
		}

	case crc_botInlineMessageMediaAuto:
		flags := m.Flags()
		_ = flags
		r = TL_botInlineMessageMediaAuto{
			flags,
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case crc_botInlineMessageText:
		flags := m.Flags()
		_ = flags
		r = TL_botInlineMessageText{
			flags,
			m.String(),
			m.FlaggedVector(flags, 1),
			m.FlaggedObject(flags, 2),
		}

	case crc_botInlineResult:
		flags := m.Flags()
		_ = flags
		r = TL_botInlineResult{
			flags,
			m.String(),
			m.String(),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedString(flags, 4),
			m.FlaggedString(flags, 5),
			m.FlaggedString(flags, 5),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 7),
			m.Object(),
		}

	case crc_messages_botResults:
		flags := m.Flags()
		_ = flags
		r = TL_messages_botResults{
			flags,
			m.Long(),
			m.FlaggedString(flags, 1),
			m.FlaggedObject(flags, 2),
			m.Vector(),
			m.Int(),
		}

	case crc_inputMessagesFilterVoice:
		r = TL_inputMessagesFilterVoice{}

	case crc_inputMessagesFilterMusic:
		r = TL_inputMessagesFilterMusic{}

	case crc_updateBotInlineSend:
		flags := m.Flags()
		_ = flags
		r = TL_updateBotInlineSend{
			flags,
			m.Int(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.String(),
			m.FlaggedObject(flags, 1),
		}

	case crc_inputPrivacyKeyChatInvite:
		r = TL_inputPrivacyKeyChatInvite{}

	case crc_privacyKeyChatInvite:
		r = TL_privacyKeyChatInvite{}

	case crc_updateEditChannelMessage:
		r = TL_updateEditChannelMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_exportedMessageLink:
		r = TL_exportedMessageLink{
			m.String(),
		}

	case crc_messageFwdHeader:
		flags := m.Flags()
		_ = flags
		r = TL_messageFwdHeader{
			flags,
			m.FlaggedInt(flags, 0),
			m.Int(),
			m.FlaggedInt(flags, 1),
			m.FlaggedInt(flags, 2),
			m.FlaggedString(flags, 3),
		}

	case crc_messageActionPinMessage:
		r = TL_messageActionPinMessage{}

	case crc_peerSettings:
		flags := m.Flags()
		_ = flags
		r = TL_peerSettings{
			flags,
		}

	case crc_updateChannelPinnedMessage:
		r = TL_updateChannelPinnedMessage{
			m.Int(),
			m.Int(),
		}

	case crc_keyboardButtonUrl:
		r = TL_keyboardButtonUrl{
			m.String(),
			m.String(),
		}

	case crc_keyboardButtonCallback:
		r = TL_keyboardButtonCallback{
			m.String(),
			m.StringBytes(),
		}

	case crc_keyboardButtonRequestPhone:
		r = TL_keyboardButtonRequestPhone{
			m.String(),
		}

	case crc_keyboardButtonRequestGeoLocation:
		r = TL_keyboardButtonRequestGeoLocation{
			m.String(),
		}

	case crc_auth_codeTypeSms:
		r = TL_auth_codeTypeSms{}

	case crc_auth_codeTypeCall:
		r = TL_auth_codeTypeCall{}

	case crc_auth_codeTypeFlashCall:
		r = TL_auth_codeTypeFlashCall{}

	case crc_auth_sentCodeTypeApp:
		r = TL_auth_sentCodeTypeApp{
			m.Int(),
		}

	case crc_auth_sentCodeTypeSms:
		r = TL_auth_sentCodeTypeSms{
			m.Int(),
		}

	case crc_auth_sentCodeTypeCall:
		r = TL_auth_sentCodeTypeCall{
			m.Int(),
		}

	case crc_auth_sentCodeTypeFlashCall:
		r = TL_auth_sentCodeTypeFlashCall{
			m.String(),
		}

	case crc_keyboardButtonSwitchInline:
		flags := m.Flags()
		_ = flags
		r = TL_keyboardButtonSwitchInline{
			flags,
			m.String(),
			m.String(),
		}

	case crc_replyInlineMarkup:
		r = TL_replyInlineMarkup{
			m.Vector(),
		}

	case crc_messages_botCallbackAnswer:
		flags := m.Flags()
		_ = flags
		r = TL_messages_botCallbackAnswer{
			flags,
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 2),
			m.Int(),
		}

	case crc_updateBotCallbackQuery:
		flags := m.Flags()
		_ = flags
		r = TL_updateBotCallbackQuery{
			flags,
			m.Long(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Long(),
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedString(flags, 1),
		}

	case crc_messages_messageEditData:
		flags := m.Flags()
		_ = flags
		r = TL_messages_messageEditData{
			flags,
		}

	case crc_updateEditMessage:
		r = TL_updateEditMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_inputBotInlineMessageMediaGeo:
		flags := m.Flags()
		_ = flags
		r = TL_inputBotInlineMessageMediaGeo{
			flags,
			m.Object(),
			m.FlaggedObject(flags, 2),
		}

	case crc_inputBotInlineMessageMediaVenue:
		flags := m.Flags()
		_ = flags
		r = TL_inputBotInlineMessageMediaVenue{
			flags,
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case crc_inputBotInlineMessageMediaContact:
		flags := m.Flags()
		_ = flags
		r = TL_inputBotInlineMessageMediaContact{
			flags,
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case crc_botInlineMessageMediaGeo:
		flags := m.Flags()
		_ = flags
		r = TL_botInlineMessageMediaGeo{
			flags,
			m.Object(),
			m.FlaggedObject(flags, 2),
		}

	case crc_botInlineMessageMediaVenue:
		flags := m.Flags()
		_ = flags
		r = TL_botInlineMessageMediaVenue{
			flags,
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case crc_botInlineMessageMediaContact:
		flags := m.Flags()
		_ = flags
		r = TL_botInlineMessageMediaContact{
			flags,
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case crc_inputBotInlineResultPhoto:
		r = TL_inputBotInlineResultPhoto{
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_inputBotInlineResultDocument:
		flags := m.Flags()
		_ = flags
		r = TL_inputBotInlineResultDocument{
			flags,
			m.String(),
			m.String(),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.Object(),
			m.Object(),
		}

	case crc_botInlineMediaResult:
		flags := m.Flags()
		_ = flags
		r = TL_botInlineMediaResult{
			flags,
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.Object(),
		}

	case crc_inputBotInlineMessageID:
		r = TL_inputBotInlineMessageID{
			m.Int(),
			m.Long(),
			m.Long(),
		}

	case crc_updateInlineBotCallbackQuery:
		flags := m.Flags()
		_ = flags
		r = TL_updateInlineBotCallbackQuery{
			flags,
			m.Long(),
			m.Int(),
			m.Object(),
			m.Long(),
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedString(flags, 1),
		}

	case crc_inlineBotSwitchPM:
		r = TL_inlineBotSwitchPM{
			m.String(),
			m.String(),
		}

	case crc_messageEntityMentionName:
		r = TL_messageEntityMentionName{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputMessageEntityMentionName:
		r = TL_inputMessageEntityMentionName{
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_messages_peerDialogs:
		r = TL_messages_peerDialogs{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crc_topPeer:
		r = TL_topPeer{
			m.Object(),
			m.Double(),
		}

	case crc_topPeerCategoryBotsPM:
		r = TL_topPeerCategoryBotsPM{}

	case crc_topPeerCategoryBotsInline:
		r = TL_topPeerCategoryBotsInline{}

	case crc_topPeerCategoryCorrespondents:
		r = TL_topPeerCategoryCorrespondents{}

	case crc_topPeerCategoryGroups:
		r = TL_topPeerCategoryGroups{}

	case crc_topPeerCategoryChannels:
		r = TL_topPeerCategoryChannels{}

	case crc_topPeerCategoryPeers:
		r = TL_topPeerCategoryPeers{
			m.Object(),
			m.Int(),
			m.Vector(),
		}

	case crc_contacts_topPeersNotModified:
		r = TL_contacts_topPeersNotModified{}

	case crc_contacts_topPeers:
		r = TL_contacts_topPeers{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_inputMessagesFilterChatPhotos:
		r = TL_inputMessagesFilterChatPhotos{}

	case crc_updateReadChannelOutbox:
		r = TL_updateReadChannelOutbox{
			m.Int(),
			m.Int(),
		}

	case crc_updateDraftMessage:
		r = TL_updateDraftMessage{
			m.Object(),
			m.Object(),
		}

	case crc_draftMessageEmpty:
		r = TL_draftMessageEmpty{}

	case crc_draftMessage:
		flags := m.Flags()
		_ = flags
		r = TL_draftMessage{
			flags,
			m.FlaggedInt(flags, 0),
			m.String(),
			m.FlaggedVector(flags, 3),
			m.Int(),
		}

	case crc_messageActionHistoryClear:
		r = TL_messageActionHistoryClear{}

	case crc_updateReadFeaturedStickers:
		r = TL_updateReadFeaturedStickers{}

	case crc_updateRecentStickers:
		r = TL_updateRecentStickers{}

	case crc_messages_featuredStickersNotModified:
		r = TL_messages_featuredStickersNotModified{}

	case crc_messages_featuredStickers:
		r = TL_messages_featuredStickers{
			m.Int(),
			m.Vector(),
			m.VectorLong(),
		}

	case crc_messages_recentStickersNotModified:
		r = TL_messages_recentStickersNotModified{}

	case crc_messages_recentStickers:
		r = TL_messages_recentStickers{
			m.Int(),
			m.Vector(),
		}

	case crc_messages_archivedStickers:
		r = TL_messages_archivedStickers{
			m.Int(),
			m.Vector(),
		}

	case crc_messages_stickerSetInstallResultSuccess:
		r = TL_messages_stickerSetInstallResultSuccess{}

	case crc_messages_stickerSetInstallResultArchive:
		r = TL_messages_stickerSetInstallResultArchive{
			m.Vector(),
		}

	case crc_stickerSetCovered:
		r = TL_stickerSetCovered{
			m.Object(),
			m.Object(),
		}

	case crc_inputMediaPhotoExternal:
		flags := m.Flags()
		_ = flags
		r = TL_inputMediaPhotoExternal{
			flags,
			m.String(),
			m.String(),
			m.FlaggedInt(flags, 0),
		}

	case crc_inputMediaDocumentExternal:
		flags := m.Flags()
		_ = flags
		r = TL_inputMediaDocumentExternal{
			flags,
			m.String(),
			m.String(),
			m.FlaggedInt(flags, 0),
		}

	case crc_updateConfig:
		r = TL_updateConfig{}

	case crc_updatePtsChanged:
		r = TL_updatePtsChanged{}

	case crc_messageActionGameScore:
		r = TL_messageActionGameScore{
			m.Long(),
			m.Int(),
		}

	case crc_documentAttributeHasStickers:
		r = TL_documentAttributeHasStickers{}

	case crc_keyboardButtonGame:
		r = TL_keyboardButtonGame{
			m.String(),
		}

	case crc_stickerSetMultiCovered:
		r = TL_stickerSetMultiCovered{
			m.Object(),
			m.Vector(),
		}

	case crc_maskCoords:
		r = TL_maskCoords{
			m.Int(),
			m.Double(),
			m.Double(),
			m.Double(),
		}

	case crc_inputStickeredMediaPhoto:
		r = TL_inputStickeredMediaPhoto{
			m.Object(),
		}

	case crc_inputStickeredMediaDocument:
		r = TL_inputStickeredMediaDocument{
			m.Object(),
		}

	case crc_inputMediaGame:
		r = TL_inputMediaGame{
			m.Object(),
		}

	case crc_messageMediaGame:
		r = TL_messageMediaGame{
			m.Object(),
		}

	case crc_inputBotInlineMessageGame:
		flags := m.Flags()
		_ = flags
		r = TL_inputBotInlineMessageGame{
			flags,
			m.FlaggedObject(flags, 2),
		}

	case crc_inputBotInlineResultGame:
		r = TL_inputBotInlineResultGame{
			m.String(),
			m.String(),
			m.Object(),
		}

	case crc_game:
		flags := m.Flags()
		_ = flags
		r = TL_game{
			flags,
			m.Long(),
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
			m.FlaggedObject(flags, 0),
		}

	case crc_inputGameID:
		r = TL_inputGameID{
			m.Long(),
			m.Long(),
		}

	case crc_inputGameShortName:
		r = TL_inputGameShortName{
			m.Object(),
			m.String(),
		}

	case crc_highScore:
		r = TL_highScore{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_highScores:
		r = TL_messages_highScores{
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_chatsSlice:
		r = TL_messages_chatsSlice{
			m.Int(),
			m.Vector(),
		}

	case crc_updateChannelWebPage:
		r = TL_updateChannelWebPage{
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_updates_differenceTooLong:
		r = TL_updates_differenceTooLong{
			m.Int(),
		}

	case crc_sendMessageGamePlayAction:
		r = TL_sendMessageGamePlayAction{}

	case crc_webPageNotModified:
		r = TL_webPageNotModified{}

	case crc_textEmpty:
		r = TL_textEmpty{}

	case crc_textPlain:
		r = TL_textPlain{
			m.String(),
		}

	case crc_textBold:
		r = TL_textBold{
			m.Object(),
		}

	case crc_textItalic:
		r = TL_textItalic{
			m.Object(),
		}

	case crc_textUnderline:
		r = TL_textUnderline{
			m.Object(),
		}

	case crc_textStrike:
		r = TL_textStrike{
			m.Object(),
		}

	case crc_textFixed:
		r = TL_textFixed{
			m.Object(),
		}

	case crc_textUrl:
		r = TL_textUrl{
			m.Object(),
			m.String(),
			m.Long(),
		}

	case crc_textEmail:
		r = TL_textEmail{
			m.Object(),
			m.String(),
		}

	case crc_textConcat:
		r = TL_textConcat{
			m.Vector(),
		}

	case crc_pageBlockTitle:
		r = TL_pageBlockTitle{
			m.Object(),
		}

	case crc_pageBlockSubtitle:
		r = TL_pageBlockSubtitle{
			m.Object(),
		}

	case crc_pageBlockAuthorDate:
		r = TL_pageBlockAuthorDate{
			m.Object(),
			m.Int(),
		}

	case crc_pageBlockHeader:
		r = TL_pageBlockHeader{
			m.Object(),
		}

	case crc_pageBlockSubheader:
		r = TL_pageBlockSubheader{
			m.Object(),
		}

	case crc_pageBlockParagraph:
		r = TL_pageBlockParagraph{
			m.Object(),
		}

	case crc_pageBlockPreformatted:
		r = TL_pageBlockPreformatted{
			m.Object(),
			m.String(),
		}

	case crc_pageBlockFooter:
		r = TL_pageBlockFooter{
			m.Object(),
		}

	case crc_pageBlockDivider:
		r = TL_pageBlockDivider{}

	case crc_pageBlockList:
		r = TL_pageBlockList{
			m.Object(),
			m.Vector(),
		}

	case crc_pageBlockBlockquote:
		r = TL_pageBlockBlockquote{
			m.Object(),
			m.Object(),
		}

	case crc_pageBlockPullquote:
		r = TL_pageBlockPullquote{
			m.Object(),
			m.Object(),
		}

	case crc_pageBlockPhoto:
		r = TL_pageBlockPhoto{
			m.Long(),
			m.Object(),
		}

	case crc_pageBlockVideo:
		flags := m.Flags()
		_ = flags
		r = TL_pageBlockVideo{
			flags,
			m.Long(),
			m.Object(),
		}

	case crc_pageBlockCover:
		r = TL_pageBlockCover{
			m.Object(),
		}

	case crc_pageBlockEmbed:
		flags := m.Flags()
		_ = flags
		r = TL_pageBlockEmbed{
			flags,
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedLong(flags, 4),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_pageBlockEmbedPost:
		r = TL_pageBlockEmbedPost{
			m.String(),
			m.Long(),
			m.Long(),
			m.String(),
			m.Int(),
			m.Vector(),
			m.Object(),
		}

	case crc_pageBlockSlideshow:
		r = TL_pageBlockSlideshow{
			m.Vector(),
			m.Object(),
		}

	case crc_pagePart:
		r = TL_pagePart{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_pageFull:
		r = TL_pageFull{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_updatePhoneCall:
		r = TL_updatePhoneCall{
			m.Object(),
		}

	case crc_updateDialogPinned:
		flags := m.Flags()
		_ = flags
		r = TL_updateDialogPinned{
			flags,
			m.Object(),
		}

	case crc_updatePinnedDialogs:
		flags := m.Flags()
		_ = flags
		r = TL_updatePinnedDialogs{
			flags,
			m.FlaggedVector(flags, 0),
		}

	case crc_inputPrivacyKeyPhoneCall:
		r = TL_inputPrivacyKeyPhoneCall{}

	case crc_privacyKeyPhoneCall:
		r = TL_privacyKeyPhoneCall{}

	case crc_pageBlockUnsupported:
		r = TL_pageBlockUnsupported{}

	case crc_pageBlockAnchor:
		r = TL_pageBlockAnchor{
			m.String(),
		}

	case crc_pageBlockCollage:
		r = TL_pageBlockCollage{
			m.Vector(),
			m.Object(),
		}

	case crc_inputPhoneCall:
		r = TL_inputPhoneCall{
			m.Long(),
			m.Long(),
		}

	case crc_phoneCallEmpty:
		r = TL_phoneCallEmpty{
			m.Long(),
		}

	case crc_phoneCallWaiting:
		flags := m.Flags()
		_ = flags
		r = TL_phoneCallWaiting{
			flags,
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.FlaggedInt(flags, 0),
		}

	case crc_phoneCallRequested:
		r = TL_phoneCallRequested{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_phoneCall:
		r = TL_phoneCall{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Long(),
			m.Object(),
			m.Object(),
			m.Vector(),
			m.Int(),
		}

	case crc_phoneCallDiscarded:
		flags := m.Flags()
		_ = flags
		r = TL_phoneCallDiscarded{
			flags,
			m.Long(),
			m.FlaggedObject(flags, 0),
			m.FlaggedInt(flags, 1),
		}

	case crc_phoneConnection:
		r = TL_phoneConnection{
			m.Long(),
			m.String(),
			m.String(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_phoneCallProtocol:
		flags := m.Flags()
		_ = flags
		r = TL_phoneCallProtocol{
			flags,
			m.Int(),
			m.Int(),
		}

	case crc_phone_phoneCall:
		r = TL_phone_phoneCall{
			m.Object(),
			m.Vector(),
		}

	case crc_phoneCallDiscardReasonMissed:
		r = TL_phoneCallDiscardReasonMissed{}

	case crc_phoneCallDiscardReasonDisconnect:
		r = TL_phoneCallDiscardReasonDisconnect{}

	case crc_phoneCallDiscardReasonHangup:
		r = TL_phoneCallDiscardReasonHangup{}

	case crc_phoneCallDiscardReasonBusy:
		r = TL_phoneCallDiscardReasonBusy{}

	case crc_inputMessagesFilterPhoneCalls:
		flags := m.Flags()
		_ = flags
		r = TL_inputMessagesFilterPhoneCalls{
			flags,
		}

	case crc_messageActionPhoneCall:
		flags := m.Flags()
		_ = flags
		r = TL_messageActionPhoneCall{
			flags,
			m.Long(),
			m.FlaggedObject(flags, 0),
			m.FlaggedInt(flags, 1),
		}

	case crc_invoice:
		flags := m.Flags()
		_ = flags
		r = TL_invoice{
			flags,
			m.String(),
			m.Vector(),
		}

	case crc_inputMediaInvoice:
		flags := m.Flags()
		_ = flags
		r = TL_inputMediaInvoice{
			flags,
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.Object(),
			m.StringBytes(),
			m.String(),
			m.String(),
		}

	case crc_messageActionPaymentSentMe:
		flags := m.Flags()
		_ = flags
		r = TL_messageActionPaymentSentMe{
			flags,
			m.String(),
			m.Long(),
			m.StringBytes(),
			m.FlaggedObject(flags, 0),
			m.FlaggedString(flags, 1),
			m.Object(),
		}

	case crc_messageMediaInvoice:
		flags := m.Flags()
		_ = flags
		r = TL_messageMediaInvoice{
			flags,
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.FlaggedInt(flags, 2),
			m.String(),
			m.Long(),
			m.String(),
		}

	case crc_keyboardButtonBuy:
		r = TL_keyboardButtonBuy{
			m.String(),
		}

	case crc_messageActionPaymentSent:
		r = TL_messageActionPaymentSent{
			m.String(),
			m.Long(),
		}

	case crc_payments_paymentForm:
		flags := m.Flags()
		_ = flags
		r = TL_payments_paymentForm{
			flags,
			m.Int(),
			m.Object(),
			m.Int(),
			m.String(),
			m.FlaggedString(flags, 4),
			m.FlaggedObject(flags, 4),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.Vector(),
		}

	case crc_postAddress:
		r = TL_postAddress{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_paymentRequestedInfo:
		flags := m.Flags()
		_ = flags
		r = TL_paymentRequestedInfo{
			flags,
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedObject(flags, 3),
		}

	case crc_updateBotWebhookJSON:
		r = TL_updateBotWebhookJSON{
			m.Object(),
		}

	case crc_updateBotWebhookJSONQuery:
		r = TL_updateBotWebhookJSONQuery{
			m.Long(),
			m.Object(),
			m.Int(),
		}

	case crc_updateBotShippingQuery:
		r = TL_updateBotShippingQuery{
			m.Long(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_updateBotPrecheckoutQuery:
		flags := m.Flags()
		_ = flags
		r = TL_updateBotPrecheckoutQuery{
			flags,
			m.Long(),
			m.Int(),
			m.StringBytes(),
			m.FlaggedObject(flags, 0),
			m.FlaggedString(flags, 1),
			m.String(),
			m.Long(),
		}

	case crc_dataJSON:
		r = TL_dataJSON{
			m.String(),
		}

	case crc_labeledPrice:
		r = TL_labeledPrice{
			m.String(),
			m.Long(),
		}

	case crc_paymentCharge:
		r = TL_paymentCharge{
			m.String(),
			m.String(),
		}

	case crc_paymentSavedCredentialsCard:
		r = TL_paymentSavedCredentialsCard{
			m.String(),
			m.String(),
		}

	case crc_webDocument:
		r = TL_webDocument{
			m.String(),
			m.Long(),
			m.Int(),
			m.String(),
			m.Vector(),
			m.Int(),
		}

	case crc_inputWebDocument:
		r = TL_inputWebDocument{
			m.String(),
			m.Int(),
			m.String(),
			m.Vector(),
		}

	case crc_inputWebFileLocation:
		r = TL_inputWebFileLocation{
			m.String(),
			m.Long(),
		}

	case crc_upload_webFile:
		r = TL_upload_webFile{
			m.Int(),
			m.String(),
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_payments_validatedRequestedInfo:
		flags := m.Flags()
		_ = flags
		r = TL_payments_validatedRequestedInfo{
			flags,
			m.FlaggedString(flags, 0),
			m.FlaggedVector(flags, 1),
		}

	case crc_payments_paymentResult:
		r = TL_payments_paymentResult{
			m.Object(),
		}

	case crc_payments_paymentVerficationNeeded:
		r = TL_payments_paymentVerficationNeeded{
			m.String(),
		}

	case crc_payments_paymentReceipt:
		flags := m.Flags()
		_ = flags
		r = TL_payments_paymentReceipt{
			flags,
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.String(),
			m.Long(),
			m.String(),
			m.Vector(),
		}

	case crc_payments_savedInfo:
		flags := m.Flags()
		_ = flags
		r = TL_payments_savedInfo{
			flags,
			m.FlaggedObject(flags, 0),
		}

	case crc_inputPaymentCredentialsSaved:
		r = TL_inputPaymentCredentialsSaved{
			m.String(),
			m.StringBytes(),
		}

	case crc_inputPaymentCredentials:
		flags := m.Flags()
		_ = flags
		r = TL_inputPaymentCredentials{
			flags,
			m.Object(),
		}

	case crc_account_tmpPassword:
		r = TL_account_tmpPassword{
			m.StringBytes(),
			m.Int(),
		}

	case crc_shippingOption:
		r = TL_shippingOption{
			m.String(),
			m.String(),
			m.Vector(),
		}

	case crc_phoneCallAccepted:
		r = TL_phoneCallAccepted{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_inputMessagesFilterRoundVoice:
		r = TL_inputMessagesFilterRoundVoice{}

	case crc_inputMessagesFilterRoundVideo:
		r = TL_inputMessagesFilterRoundVideo{}

	case crc_upload_fileCdnRedirect:
		r = TL_upload_fileCdnRedirect{
			m.Int(),
			m.StringBytes(),
			m.StringBytes(),
			m.StringBytes(),
			m.Vector(),
		}

	case crc_sendMessageRecordRoundAction:
		r = TL_sendMessageRecordRoundAction{}

	case crc_sendMessageUploadRoundAction:
		r = TL_sendMessageUploadRoundAction{
			m.Int(),
		}

	case crc_upload_cdnFileReuploadNeeded:
		r = TL_upload_cdnFileReuploadNeeded{
			m.StringBytes(),
		}

	case crc_upload_cdnFile:
		r = TL_upload_cdnFile{
			m.StringBytes(),
		}

	case crc_cdnPublicKey:
		r = TL_cdnPublicKey{
			m.Int(),
			m.String(),
		}

	case crc_cdnConfig:
		r = TL_cdnConfig{
			m.Vector(),
		}

	case crc_updateLangPackTooLong:
		r = TL_updateLangPackTooLong{}

	case crc_updateLangPack:
		r = TL_updateLangPack{
			m.Object(),
		}

	case crc_pageBlockChannel:
		r = TL_pageBlockChannel{
			m.Object(),
		}

	case crc_inputStickerSetItem:
		flags := m.Flags()
		_ = flags
		r = TL_inputStickerSetItem{
			flags,
			m.Object(),
			m.String(),
			m.FlaggedObject(flags, 0),
		}

	case crc_langPackString:
		r = TL_langPackString{
			m.String(),
			m.String(),
		}

	case crc_langPackStringPluralized:
		flags := m.Flags()
		_ = flags
		r = TL_langPackStringPluralized{
			flags,
			m.String(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedString(flags, 4),
			m.String(),
		}

	case crc_langPackStringDeleted:
		r = TL_langPackStringDeleted{
			m.String(),
		}

	case crc_langPackDifference:
		r = TL_langPackDifference{
			m.String(),
			m.Int(),
			m.Int(),
			m.Vector(),
		}

	case crc_langPackLanguage:
		r = TL_langPackLanguage{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_channelParticipantAdmin:
		flags := m.Flags()
		_ = flags
		r = TL_channelParticipantAdmin{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_channelParticipantBanned:
		flags := m.Flags()
		_ = flags
		r = TL_channelParticipantBanned{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_channelParticipantsBanned:
		r = TL_channelParticipantsBanned{
			m.String(),
		}

	case crc_channelParticipantsSearch:
		r = TL_channelParticipantsSearch{
			m.String(),
		}

	case crc_topPeerCategoryPhoneCalls:
		r = TL_topPeerCategoryPhoneCalls{}

	case crc_pageBlockAudio:
		r = TL_pageBlockAudio{
			m.Long(),
			m.Object(),
		}

	case crc_channelAdminRights:
		flags := m.Flags()
		_ = flags
		r = TL_channelAdminRights{
			flags,
		}

	case crc_channelBannedRights:
		flags := m.Flags()
		_ = flags
		r = TL_channelBannedRights{
			flags,
			m.Int(),
		}

	case crc_channelAdminLogEventActionChangeTitle:
		r = TL_channelAdminLogEventActionChangeTitle{
			m.String(),
			m.String(),
		}

	case crc_channelAdminLogEventActionChangeAbout:
		r = TL_channelAdminLogEventActionChangeAbout{
			m.String(),
			m.String(),
		}

	case crc_channelAdminLogEventActionChangeUsername:
		r = TL_channelAdminLogEventActionChangeUsername{
			m.String(),
			m.String(),
		}

	case crc_channelAdminLogEventActionChangePhoto:
		r = TL_channelAdminLogEventActionChangePhoto{
			m.Object(),
			m.Object(),
		}

	case crc_channelAdminLogEventActionToggleInvites:
		r = TL_channelAdminLogEventActionToggleInvites{
			m.Object(),
		}

	case crc_channelAdminLogEventActionToggleSignatures:
		r = TL_channelAdminLogEventActionToggleSignatures{
			m.Object(),
		}

	case crc_channelAdminLogEventActionUpdatePinned:
		r = TL_channelAdminLogEventActionUpdatePinned{
			m.Object(),
		}

	case crc_channelAdminLogEventActionEditMessage:
		r = TL_channelAdminLogEventActionEditMessage{
			m.Object(),
			m.Object(),
		}

	case crc_channelAdminLogEventActionDeleteMessage:
		r = TL_channelAdminLogEventActionDeleteMessage{
			m.Object(),
		}

	case crc_channelAdminLogEventActionParticipantJoin:
		r = TL_channelAdminLogEventActionParticipantJoin{}

	case crc_channelAdminLogEventActionParticipantLeave:
		r = TL_channelAdminLogEventActionParticipantLeave{}

	case crc_channelAdminLogEventActionParticipantInvite:
		r = TL_channelAdminLogEventActionParticipantInvite{
			m.Object(),
		}

	case crc_channelAdminLogEventActionParticipantToggleBan:
		r = TL_channelAdminLogEventActionParticipantToggleBan{
			m.Object(),
			m.Object(),
		}

	case crc_channelAdminLogEventActionParticipantToggleAdmin:
		r = TL_channelAdminLogEventActionParticipantToggleAdmin{
			m.Object(),
			m.Object(),
		}

	case crc_channelAdminLogEvent:
		r = TL_channelAdminLogEvent{
			m.Long(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_channels_adminLogResults:
		r = TL_channels_adminLogResults{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_channelAdminLogEventsFilter:
		flags := m.Flags()
		_ = flags
		r = TL_channelAdminLogEventsFilter{
			flags,
		}

	case crc_messageActionScreenshotTaken:
		r = TL_messageActionScreenshotTaken{}

	case crc_popularContact:
		r = TL_popularContact{
			m.Long(),
			m.Int(),
		}

	case crc_cdnFileHash:
		r = TL_cdnFileHash{
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_inputMessagesFilterMyMentions:
		r = TL_inputMessagesFilterMyMentions{}

	case crc_inputMessagesFilterMyMentionsUnread:
		r = TL_inputMessagesFilterMyMentionsUnread{}

	case crc_updateContactsReset:
		r = TL_updateContactsReset{}

	case crc_channelAdminLogEventActionChangeStickerSet:
		r = TL_channelAdminLogEventActionChangeStickerSet{
			m.Object(),
			m.Object(),
		}

	case crc_updateFavedStickers:
		r = TL_updateFavedStickers{}

	case crc_messages_favedStickers:
		r = TL_messages_favedStickers{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_favedStickersNotModified:
		r = TL_messages_favedStickersNotModified{}

	case crc_updateChannelReadMessagesContents:
		r = TL_updateChannelReadMessagesContents{
			m.Int(),
			m.VectorInt(),
		}

	case crc_invokeAfterMsg:
		r = TL_invokeAfterMsg{
			m.Long(),
			m.Object(),
		}

	case crc_invokeAfterMsgs:
		r = TL_invokeAfterMsgs{
			m.VectorLong(),
			m.Object(),
		}

	case crc_auth_checkPhone:
		r = TL_auth_checkPhone{
			m.String(),
		}

	case crc_auth_sendCode:
		flags := m.Flags()
		_ = flags
		r = TL_auth_sendCode{
			flags,
			m.String(),
			m.FlaggedObject(flags, 0),
			m.Int(),
			m.String(),
		}

	case crc_auth_signUp:
		r = TL_auth_signUp{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_auth_signIn:
		r = TL_auth_signIn{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_auth_logOut:
		r = TL_auth_logOut{}

	case crc_auth_resetAuthorizations:
		r = TL_auth_resetAuthorizations{}

	case crc_auth_sendInvites:
		r = TL_auth_sendInvites{
			m.VectorString(),
			m.String(),
		}

	case crc_auth_exportAuthorization:
		r = TL_auth_exportAuthorization{
			m.Int(),
		}

	case crc_auth_importAuthorization:
		r = TL_auth_importAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case crc_account_registerDevice:
		r = TL_account_registerDevice{
			m.Int(),
			m.String(),
		}

	case crc_account_unregisterDevice:
		r = TL_account_unregisterDevice{
			m.Int(),
			m.String(),
		}

	case crc_account_updateNotifySettings:
		r = TL_account_updateNotifySettings{
			m.Object(),
			m.Object(),
		}

	case crc_account_getNotifySettings:
		r = TL_account_getNotifySettings{
			m.Object(),
		}

	case crc_account_resetNotifySettings:
		r = TL_account_resetNotifySettings{}

	case crc_account_updateProfile:
		flags := m.Flags()
		_ = flags
		r = TL_account_updateProfile{
			flags,
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
		}

	case crc_account_updateStatus:
		r = TL_account_updateStatus{
			m.Object(),
		}

	case crc_account_getWallPapers:
		r = TL_account_getWallPapers{}

	case crc_users_getUsers:
		r = TL_users_getUsers{
			m.Vector(),
		}

	case crc_users_getFullUser:
		r = TL_users_getFullUser{
			m.Object(),
		}

	case crc_contacts_getStatuses:
		r = TL_contacts_getStatuses{}

	case crc_contacts_getContacts:
		r = TL_contacts_getContacts{
			m.Int(),
		}

	case crc_contacts_importContacts:
		r = TL_contacts_importContacts{
			m.Vector(),
		}

	case crc_contacts_search:
		r = TL_contacts_search{
			m.String(),
			m.Int(),
		}

	case crc_contacts_deleteContact:
		r = TL_contacts_deleteContact{
			m.Object(),
		}

	case crc_contacts_deleteContacts:
		r = TL_contacts_deleteContacts{
			m.Vector(),
		}

	case crc_contacts_block:
		r = TL_contacts_block{
			m.Object(),
		}

	case crc_contacts_unblock:
		r = TL_contacts_unblock{
			m.Object(),
		}

	case crc_contacts_getBlocked:
		r = TL_contacts_getBlocked{
			m.Int(),
			m.Int(),
		}

	case crc_messages_getMessages:
		r = TL_messages_getMessages{
			m.VectorInt(),
		}

	case crc_messages_getDialogs:
		flags := m.Flags()
		_ = flags
		r = TL_messages_getDialogs{
			flags,
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_messages_getHistory:
		r = TL_messages_getHistory{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_search:
		flags := m.Flags()
		_ = flags
		r = TL_messages_search{
			flags,
			m.Object(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_readHistory:
		r = TL_messages_readHistory{
			m.Object(),
			m.Int(),
		}

	case crc_messages_deleteHistory:
		flags := m.Flags()
		_ = flags
		r = TL_messages_deleteHistory{
			flags,
			m.Object(),
			m.Int(),
		}

	case crc_messages_deleteMessages:
		flags := m.Flags()
		_ = flags
		r = TL_messages_deleteMessages{
			flags,
			m.VectorInt(),
		}

	case crc_messages_receivedMessages:
		r = TL_messages_receivedMessages{
			m.Int(),
		}

	case crc_messages_setTyping:
		r = TL_messages_setTyping{
			m.Object(),
			m.Object(),
		}

	case crc_messages_sendMessage:
		flags := m.Flags()
		_ = flags
		r = TL_messages_sendMessage{
			flags,
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.String(),
			m.Long(),
			m.FlaggedObject(flags, 2),
			m.FlaggedVector(flags, 3),
		}

	case crc_messages_sendMedia:
		flags := m.Flags()
		_ = flags
		r = TL_messages_sendMedia{
			flags,
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.Object(),
			m.Long(),
			m.FlaggedObject(flags, 2),
		}

	case crc_messages_forwardMessages:
		flags := m.Flags()
		_ = flags
		r = TL_messages_forwardMessages{
			flags,
			m.Object(),
			m.VectorInt(),
			m.VectorLong(),
			m.Object(),
		}

	case crc_messages_getChats:
		r = TL_messages_getChats{
			m.VectorInt(),
		}

	case crc_messages_getFullChat:
		r = TL_messages_getFullChat{
			m.Int(),
		}

	case crc_messages_editChatTitle:
		r = TL_messages_editChatTitle{
			m.Int(),
			m.String(),
		}

	case crc_messages_editChatPhoto:
		r = TL_messages_editChatPhoto{
			m.Int(),
			m.Object(),
		}

	case crc_messages_addChatUser:
		r = TL_messages_addChatUser{
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_messages_deleteChatUser:
		r = TL_messages_deleteChatUser{
			m.Int(),
			m.Object(),
		}

	case crc_messages_createChat:
		r = TL_messages_createChat{
			m.Vector(),
			m.String(),
		}

	case crc_updates_getState:
		r = TL_updates_getState{}

	case crc_updates_getDifference:
		flags := m.Flags()
		_ = flags
		r = TL_updates_getDifference{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 0),
			m.Int(),
			m.Int(),
		}

	case crc_photos_updateProfilePhoto:
		r = TL_photos_updateProfilePhoto{
			m.Object(),
		}

	case crc_photos_uploadProfilePhoto:
		r = TL_photos_uploadProfilePhoto{
			m.Object(),
		}

	case crc_upload_saveFilePart:
		r = TL_upload_saveFilePart{
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_upload_getFile:
		r = TL_upload_getFile{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_help_getConfig:
		r = TL_help_getConfig{}

	case crc_help_getNearestDc:
		r = TL_help_getNearestDc{}

	case crc_help_getAppUpdate:
		r = TL_help_getAppUpdate{}

	case crc_help_saveAppLog:
		r = TL_help_saveAppLog{
			m.Vector(),
		}

	case crc_help_getInviteText:
		r = TL_help_getInviteText{}

	case crc_photos_deletePhotos:
		r = TL_photos_deletePhotos{
			m.Vector(),
		}

	case crc_photos_getUserPhotos:
		r = TL_photos_getUserPhotos{
			m.Object(),
			m.Int(),
			m.Long(),
			m.Int(),
		}

	case crc_messages_forwardMessage:
		r = TL_messages_forwardMessage{
			m.Object(),
			m.Int(),
			m.Long(),
		}

	case crc_messages_getDhConfig:
		r = TL_messages_getDhConfig{
			m.Int(),
			m.Int(),
		}

	case crc_messages_requestEncryption:
		r = TL_messages_requestEncryption{
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messages_acceptEncryption:
		r = TL_messages_acceptEncryption{
			m.Object(),
			m.StringBytes(),
			m.Long(),
		}

	case crc_messages_discardEncryption:
		r = TL_messages_discardEncryption{
			m.Int(),
		}

	case crc_messages_setEncryptedTyping:
		r = TL_messages_setEncryptedTyping{
			m.Object(),
			m.Object(),
		}

	case crc_messages_readEncryptedHistory:
		r = TL_messages_readEncryptedHistory{
			m.Object(),
			m.Int(),
		}

	case crc_messages_sendEncrypted:
		r = TL_messages_sendEncrypted{
			m.Object(),
			m.Long(),
			m.StringBytes(),
		}

	case crc_messages_sendEncryptedFile:
		r = TL_messages_sendEncryptedFile{
			m.Object(),
			m.Long(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_messages_sendEncryptedService:
		r = TL_messages_sendEncryptedService{
			m.Object(),
			m.Long(),
			m.StringBytes(),
		}

	case crc_messages_receivedQueue:
		r = TL_messages_receivedQueue{
			m.Int(),
		}

	case crc_upload_saveBigFilePart:
		r = TL_upload_saveBigFilePart{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_initConnection:
		r = TL_initConnection{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
		}

	case crc_help_getSupport:
		r = TL_help_getSupport{}

	case crc_auth_bindTempAuthKey:
		r = TL_auth_bindTempAuthKey{
			m.Long(),
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_contacts_exportCard:
		r = TL_contacts_exportCard{}

	case crc_contacts_importCard:
		r = TL_contacts_importCard{
			m.VectorInt(),
		}

	case crc_messages_readMessageContents:
		r = TL_messages_readMessageContents{
			m.VectorInt(),
		}

	case crc_account_checkUsername:
		r = TL_account_checkUsername{
			m.String(),
		}

	case crc_account_updateUsername:
		r = TL_account_updateUsername{
			m.String(),
		}

	case crc_account_getPrivacy:
		r = TL_account_getPrivacy{
			m.Object(),
		}

	case crc_account_setPrivacy:
		r = TL_account_setPrivacy{
			m.Object(),
			m.Vector(),
		}

	case crc_account_deleteAccount:
		r = TL_account_deleteAccount{
			m.String(),
		}

	case crc_account_getAccountTTL:
		r = TL_account_getAccountTTL{}

	case crc_account_setAccountTTL:
		r = TL_account_setAccountTTL{
			m.Object(),
		}

	case crc_invokeWithLayer:
		r = TL_invokeWithLayer{
			m.Int(),
			m.Object(),
		}

	case crc_contacts_resolveUsername:
		r = TL_contacts_resolveUsername{
			m.String(),
		}

	case crc_account_sendChangePhoneCode:
		flags := m.Flags()
		_ = flags
		r = TL_account_sendChangePhoneCode{
			flags,
			m.String(),
			m.FlaggedObject(flags, 0),
		}

	case crc_account_changePhone:
		r = TL_account_changePhone{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_messages_getAllStickers:
		r = TL_messages_getAllStickers{
			m.Int(),
		}

	case crc_account_updateDeviceLocked:
		r = TL_account_updateDeviceLocked{
			m.Int(),
		}

	case crc_account_getPassword:
		r = TL_account_getPassword{}

	case crc_auth_checkPassword:
		r = TL_auth_checkPassword{
			m.StringBytes(),
		}

	case crc_messages_getWebPagePreview:
		r = TL_messages_getWebPagePreview{
			m.String(),
		}

	case crc_account_getAuthorizations:
		r = TL_account_getAuthorizations{}

	case crc_account_resetAuthorization:
		r = TL_account_resetAuthorization{
			m.Long(),
		}

	case crc_account_getPasswordSettings:
		r = TL_account_getPasswordSettings{
			m.StringBytes(),
		}

	case crc_account_updatePasswordSettings:
		r = TL_account_updatePasswordSettings{
			m.StringBytes(),
			m.Object(),
		}

	case crc_auth_requestPasswordRecovery:
		r = TL_auth_requestPasswordRecovery{}

	case crc_auth_recoverPassword:
		r = TL_auth_recoverPassword{
			m.String(),
		}

	case crc_invokeWithoutUpdates:
		r = TL_invokeWithoutUpdates{
			m.Object(),
		}

	case crc_messages_exportChatInvite:
		r = TL_messages_exportChatInvite{
			m.Int(),
		}

	case crc_messages_checkChatInvite:
		r = TL_messages_checkChatInvite{
			m.String(),
		}

	case crc_messages_importChatInvite:
		r = TL_messages_importChatInvite{
			m.String(),
		}

	case crc_messages_getStickerSet:
		r = TL_messages_getStickerSet{
			m.Object(),
		}

	case crc_messages_installStickerSet:
		r = TL_messages_installStickerSet{
			m.Object(),
			m.Object(),
		}

	case crc_messages_uninstallStickerSet:
		r = TL_messages_uninstallStickerSet{
			m.Object(),
		}

	case crc_auth_importBotAuthorization:
		r = TL_auth_importBotAuthorization{
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crc_messages_startBot:
		r = TL_messages_startBot{
			m.Object(),
			m.Object(),
			m.Long(),
			m.String(),
		}

	case crc_help_getAppChangelog:
		r = TL_help_getAppChangelog{
			m.String(),
		}

	case crc_messages_reportSpam:
		r = TL_messages_reportSpam{
			m.Object(),
		}

	case crc_messages_getMessagesViews:
		r = TL_messages_getMessagesViews{
			m.Object(),
			m.VectorInt(),
			m.Object(),
		}

	case crc_updates_getChannelDifference:
		flags := m.Flags()
		_ = flags
		r = TL_updates_getChannelDifference{
			flags,
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_channels_readHistory:
		r = TL_channels_readHistory{
			m.Object(),
			m.Int(),
		}

	case crc_channels_deleteMessages:
		r = TL_channels_deleteMessages{
			m.Object(),
			m.VectorInt(),
		}

	case crc_channels_deleteUserHistory:
		r = TL_channels_deleteUserHistory{
			m.Object(),
			m.Object(),
		}

	case crc_channels_reportSpam:
		r = TL_channels_reportSpam{
			m.Object(),
			m.Object(),
			m.VectorInt(),
		}

	case crc_channels_getMessages:
		r = TL_channels_getMessages{
			m.Object(),
			m.VectorInt(),
		}

	case crc_channels_getParticipants:
		r = TL_channels_getParticipants{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_channels_getParticipant:
		r = TL_channels_getParticipant{
			m.Object(),
			m.Object(),
		}

	case crc_channels_getChannels:
		r = TL_channels_getChannels{
			m.Vector(),
		}

	case crc_channels_getFullChannel:
		r = TL_channels_getFullChannel{
			m.Object(),
		}

	case crc_channels_createChannel:
		flags := m.Flags()
		_ = flags
		r = TL_channels_createChannel{
			flags,
			m.String(),
			m.String(),
		}

	case crc_channels_editAbout:
		r = TL_channels_editAbout{
			m.Object(),
			m.String(),
		}

	case crc_channels_editAdmin:
		r = TL_channels_editAdmin{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_channels_editTitle:
		r = TL_channels_editTitle{
			m.Object(),
			m.String(),
		}

	case crc_channels_editPhoto:
		r = TL_channels_editPhoto{
			m.Object(),
			m.Object(),
		}

	case crc_channels_checkUsername:
		r = TL_channels_checkUsername{
			m.Object(),
			m.String(),
		}

	case crc_channels_updateUsername:
		r = TL_channels_updateUsername{
			m.Object(),
			m.String(),
		}

	case crc_channels_joinChannel:
		r = TL_channels_joinChannel{
			m.Object(),
		}

	case crc_channels_leaveChannel:
		r = TL_channels_leaveChannel{
			m.Object(),
		}

	case crc_channels_inviteToChannel:
		r = TL_channels_inviteToChannel{
			m.Object(),
			m.Vector(),
		}

	case crc_channels_exportInvite:
		r = TL_channels_exportInvite{
			m.Object(),
		}

	case crc_channels_deleteChannel:
		r = TL_channels_deleteChannel{
			m.Object(),
		}

	case crc_messages_toggleChatAdmins:
		r = TL_messages_toggleChatAdmins{
			m.Int(),
			m.Object(),
		}

	case crc_messages_editChatAdmin:
		r = TL_messages_editChatAdmin{
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_migrateChat:
		r = TL_messages_migrateChat{
			m.Int(),
		}

	case crc_messages_searchGlobal:
		r = TL_messages_searchGlobal{
			m.String(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_account_reportPeer:
		r = TL_account_reportPeer{
			m.Object(),
			m.Object(),
		}

	case crc_messages_reorderStickerSets:
		flags := m.Flags()
		_ = flags
		r = TL_messages_reorderStickerSets{
			flags,
			m.VectorLong(),
		}

	case crc_help_getTermsOfService:
		r = TL_help_getTermsOfService{}

	case crc_messages_getDocumentByHash:
		r = TL_messages_getDocumentByHash{
			m.StringBytes(),
			m.Int(),
			m.String(),
		}

	case crc_messages_searchGifs:
		r = TL_messages_searchGifs{
			m.String(),
			m.Int(),
		}

	case crc_messages_getSavedGifs:
		r = TL_messages_getSavedGifs{
			m.Int(),
		}

	case crc_messages_saveGif:
		r = TL_messages_saveGif{
			m.Object(),
			m.Object(),
		}

	case crc_messages_getInlineBotResults:
		flags := m.Flags()
		_ = flags
		r = TL_messages_getInlineBotResults{
			flags,
			m.Object(),
			m.Object(),
			m.FlaggedObject(flags, 0),
			m.String(),
			m.String(),
		}

	case crc_messages_setInlineBotResults:
		flags := m.Flags()
		_ = flags
		r = TL_messages_setInlineBotResults{
			flags,
			m.Long(),
			m.Vector(),
			m.Int(),
			m.FlaggedString(flags, 2),
			m.FlaggedObject(flags, 3),
		}

	case crc_messages_sendInlineBotResult:
		flags := m.Flags()
		_ = flags
		r = TL_messages_sendInlineBotResult{
			flags,
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.Long(),
			m.Long(),
			m.String(),
		}

	case crc_channels_toggleInvites:
		r = TL_channels_toggleInvites{
			m.Object(),
			m.Object(),
		}

	case crc_channels_exportMessageLink:
		r = TL_channels_exportMessageLink{
			m.Object(),
			m.Int(),
		}

	case crc_channels_toggleSignatures:
		r = TL_channels_toggleSignatures{
			m.Object(),
			m.Object(),
		}

	case crc_messages_hideReportSpam:
		r = TL_messages_hideReportSpam{
			m.Object(),
		}

	case crc_messages_getPeerSettings:
		r = TL_messages_getPeerSettings{
			m.Object(),
		}

	case crc_channels_updatePinnedMessage:
		flags := m.Flags()
		_ = flags
		r = TL_channels_updatePinnedMessage{
			flags,
			m.Object(),
			m.Int(),
		}

	case crc_auth_resendCode:
		r = TL_auth_resendCode{
			m.String(),
			m.String(),
		}

	case crc_auth_cancelCode:
		r = TL_auth_cancelCode{
			m.String(),
			m.String(),
		}

	case crc_messages_getMessageEditData:
		r = TL_messages_getMessageEditData{
			m.Object(),
			m.Int(),
		}

	case crc_messages_editMessage:
		flags := m.Flags()
		_ = flags
		r = TL_messages_editMessage{
			flags,
			m.Object(),
			m.Int(),
			m.FlaggedString(flags, 11),
			m.FlaggedObject(flags, 2),
			m.FlaggedVector(flags, 3),
		}

	case crc_messages_editInlineBotMessage:
		flags := m.Flags()
		_ = flags
		r = TL_messages_editInlineBotMessage{
			flags,
			m.Object(),
			m.FlaggedString(flags, 11),
			m.FlaggedObject(flags, 2),
			m.FlaggedVector(flags, 3),
		}

	case crc_messages_getBotCallbackAnswer:
		flags := m.Flags()
		_ = flags
		r = TL_messages_getBotCallbackAnswer{
			flags,
			m.Object(),
			m.Int(),
			m.FlaggedStringBytes(flags, 0),
		}

	case crc_messages_setBotCallbackAnswer:
		flags := m.Flags()
		_ = flags
		r = TL_messages_setBotCallbackAnswer{
			flags,
			m.Long(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 2),
			m.Int(),
		}

	case crc_contacts_getTopPeers:
		flags := m.Flags()
		_ = flags
		r = TL_contacts_getTopPeers{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_contacts_resetTopPeerRating:
		r = TL_contacts_resetTopPeerRating{
			m.Object(),
			m.Object(),
		}

	case crc_messages_getPeerDialogs:
		r = TL_messages_getPeerDialogs{
			m.Vector(),
		}

	case crc_messages_saveDraft:
		flags := m.Flags()
		_ = flags
		r = TL_messages_saveDraft{
			flags,
			m.FlaggedInt(flags, 0),
			m.Object(),
			m.String(),
			m.FlaggedVector(flags, 3),
		}

	case crc_messages_getAllDrafts:
		r = TL_messages_getAllDrafts{}

	case crc_account_sendConfirmPhoneCode:
		flags := m.Flags()
		_ = flags
		r = TL_account_sendConfirmPhoneCode{
			flags,
			m.String(),
			m.FlaggedObject(flags, 0),
		}

	case crc_account_confirmPhone:
		r = TL_account_confirmPhone{
			m.String(),
			m.String(),
		}

	case crc_messages_getFeaturedStickers:
		r = TL_messages_getFeaturedStickers{
			m.Int(),
		}

	case crc_messages_readFeaturedStickers:
		r = TL_messages_readFeaturedStickers{
			m.VectorLong(),
		}

	case crc_messages_getRecentStickers:
		flags := m.Flags()
		_ = flags
		r = TL_messages_getRecentStickers{
			flags,
			m.Int(),
		}

	case crc_messages_saveRecentSticker:
		flags := m.Flags()
		_ = flags
		r = TL_messages_saveRecentSticker{
			flags,
			m.Object(),
			m.Object(),
		}

	case crc_messages_clearRecentStickers:
		flags := m.Flags()
		_ = flags
		r = TL_messages_clearRecentStickers{
			flags,
		}

	case crc_messages_getArchivedStickers:
		flags := m.Flags()
		_ = flags
		r = TL_messages_getArchivedStickers{
			flags,
			m.Long(),
			m.Int(),
		}

	case crc_channels_getAdminedPublicChannels:
		r = TL_channels_getAdminedPublicChannels{}

	case crc_auth_dropTempAuthKeys:
		r = TL_auth_dropTempAuthKeys{
			m.VectorLong(),
		}

	case crc_messages_setGameScore:
		flags := m.Flags()
		_ = flags
		r = TL_messages_setGameScore{
			flags,
			m.Object(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_messages_setInlineGameScore:
		flags := m.Flags()
		_ = flags
		r = TL_messages_setInlineGameScore{
			flags,
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crc_messages_getMaskStickers:
		r = TL_messages_getMaskStickers{
			m.Int(),
		}

	case crc_messages_getAttachedStickers:
		r = TL_messages_getAttachedStickers{
			m.Object(),
		}

	case crc_messages_getGameHighScores:
		r = TL_messages_getGameHighScores{
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crc_messages_getInlineGameHighScores:
		r = TL_messages_getInlineGameHighScores{
			m.Object(),
			m.Object(),
		}

	case crc_messages_getCommonChats:
		r = TL_messages_getCommonChats{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_getAllChats:
		r = TL_messages_getAllChats{
			m.VectorInt(),
		}

	case crc_help_setBotUpdatesStatus:
		r = TL_help_setBotUpdatesStatus{
			m.Int(),
			m.String(),
		}

	case crc_messages_getWebPage:
		r = TL_messages_getWebPage{
			m.String(),
			m.Int(),
		}

	case crc_messages_toggleDialogPin:
		flags := m.Flags()
		_ = flags
		r = TL_messages_toggleDialogPin{
			flags,
			m.Object(),
		}

	case crc_messages_reorderPinnedDialogs:
		flags := m.Flags()
		_ = flags
		r = TL_messages_reorderPinnedDialogs{
			flags,
			m.Vector(),
		}

	case crc_messages_getPinnedDialogs:
		r = TL_messages_getPinnedDialogs{}

	case crc_phone_requestCall:
		r = TL_phone_requestCall{
			m.Object(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_phone_acceptCall:
		r = TL_phone_acceptCall{
			m.Object(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_phone_discardCall:
		r = TL_phone_discardCall{
			m.Object(),
			m.Int(),
			m.Object(),
			m.Long(),
		}

	case crc_phone_receivedCall:
		r = TL_phone_receivedCall{
			m.Object(),
		}

	case crc_messages_reportEncryptedSpam:
		r = TL_messages_reportEncryptedSpam{
			m.Object(),
		}

	case crc_payments_getPaymentForm:
		r = TL_payments_getPaymentForm{
			m.Int(),
		}

	case crc_payments_sendPaymentForm:
		flags := m.Flags()
		_ = flags
		r = TL_payments_sendPaymentForm{
			flags,
			m.Int(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.Object(),
		}

	case crc_account_getTmpPassword:
		r = TL_account_getTmpPassword{
			m.StringBytes(),
			m.Int(),
		}

	case crc_messages_setBotShippingResults:
		flags := m.Flags()
		_ = flags
		r = TL_messages_setBotShippingResults{
			flags,
			m.Long(),
			m.FlaggedString(flags, 0),
			m.FlaggedVector(flags, 1),
		}

	case crc_messages_setBotPrecheckoutResults:
		flags := m.Flags()
		_ = flags
		r = TL_messages_setBotPrecheckoutResults{
			flags,
			m.Long(),
			m.FlaggedString(flags, 0),
		}

	case crc_upload_getWebFile:
		r = TL_upload_getWebFile{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_bots_sendCustomRequest:
		r = TL_bots_sendCustomRequest{
			m.String(),
			m.Object(),
		}

	case crc_bots_answerWebhookJSONQuery:
		r = TL_bots_answerWebhookJSONQuery{
			m.Long(),
			m.Object(),
		}

	case crc_payments_getPaymentReceipt:
		r = TL_payments_getPaymentReceipt{
			m.Int(),
		}

	case crc_payments_validateRequestedInfo:
		flags := m.Flags()
		_ = flags
		r = TL_payments_validateRequestedInfo{
			flags,
			m.Int(),
			m.Object(),
		}

	case crc_payments_getSavedInfo:
		r = TL_payments_getSavedInfo{}

	case crc_payments_clearSavedInfo:
		flags := m.Flags()
		_ = flags
		r = TL_payments_clearSavedInfo{
			flags,
		}

	case crc_phone_getCallConfig:
		r = TL_phone_getCallConfig{}

	case crc_phone_confirmCall:
		r = TL_phone_confirmCall{
			m.Object(),
			m.StringBytes(),
			m.Long(),
			m.Object(),
		}

	case crc_phone_setCallRating:
		r = TL_phone_setCallRating{
			m.Object(),
			m.Int(),
			m.String(),
		}

	case crc_phone_saveCallDebug:
		r = TL_phone_saveCallDebug{
			m.Object(),
			m.Object(),
		}

	case crc_upload_getCdnFile:
		r = TL_upload_getCdnFile{
			m.StringBytes(),
			m.Int(),
			m.Int(),
		}

	case crc_upload_reuploadCdnFile:
		r = TL_upload_reuploadCdnFile{
			m.StringBytes(),
			m.StringBytes(),
		}

	case crc_help_getCdnConfig:
		r = TL_help_getCdnConfig{}

	case crc_messages_uploadMedia:
		r = TL_messages_uploadMedia{
			m.Object(),
			m.Object(),
		}

	case crc_stickers_createStickerSet:
		flags := m.Flags()
		_ = flags
		r = TL_stickers_createStickerSet{
			flags,
			m.Object(),
			m.String(),
			m.String(),
			m.Vector(),
		}

	case crc_langpack_getLangPack:
		r = TL_langpack_getLangPack{
			m.String(),
		}

	case crc_langpack_getStrings:
		r = TL_langpack_getStrings{
			m.String(),
			m.VectorString(),
		}

	case crc_langpack_getDifference:
		r = TL_langpack_getDifference{
			m.Int(),
		}

	case crc_langpack_getLanguages:
		r = TL_langpack_getLanguages{}

	case crc_channels_editBanned:
		r = TL_channels_editBanned{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_channels_getAdminLog:
		flags := m.Flags()
		_ = flags
		r = TL_channels_getAdminLog{
			flags,
			m.Object(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.FlaggedVector(flags, 1),
			m.Long(),
			m.Long(),
			m.Int(),
		}

	case crc_stickers_removeStickerFromSet:
		r = TL_stickers_removeStickerFromSet{
			m.Object(),
		}

	case crc_stickers_changeStickerPosition:
		r = TL_stickers_changeStickerPosition{
			m.Object(),
			m.Int(),
		}

	case crc_stickers_addStickerToSet:
		r = TL_stickers_addStickerToSet{
			m.Object(),
			m.Object(),
		}

	case crc_messages_sendScreenshotNotification:
		r = TL_messages_sendScreenshotNotification{
			m.Object(),
			m.Int(),
			m.Long(),
		}

	case crc_upload_getCdnFileHashes:
		r = TL_upload_getCdnFileHashes{
			m.StringBytes(),
			m.Int(),
		}

	case crc_messages_getUnreadMentions:
		r = TL_messages_getUnreadMentions{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_faveSticker:
		r = TL_messages_faveSticker{
			m.Object(),
			m.Object(),
		}

	case crc_channels_setStickers:
		r = TL_channels_setStickers{
			m.Object(),
			m.Object(),
		}

	case crc_contacts_resetSaved:
		r = TL_contacts_resetSaved{}

	case crc_messages_getFavedStickers:
		r = TL_messages_getFavedStickers{
			m.Int(),
		}

	case crc_channels_readMessageContents:
		r = TL_channels_readMessageContents{
			m.Object(),
			m.VectorInt(),
		}

	default:
		glog.V(3).Infoln("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		glog.V(3).Infof("Unknown constructor: %d", constructor)
		glog.V(3).Infoln("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		m.Err = fmt.Errorf("Unknown constructor: \u002508x", constructor)
		return nil

	}

	if m.Err != nil {
		glog.V(3).Infof("ObjectGenerated - constructor: %d, Error: %+v", constructor, m.Err)
		runtime.StartTrace()
		return nil
	}

	return
}
