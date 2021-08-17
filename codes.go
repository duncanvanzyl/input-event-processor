package input

/*
event codes translated directly from:
https://github.com/torvalds/linux/blob/v4.12/include/uapi/linux/input-event-codes.h
*/

/* SPDX-License-Identifier: GPL-2.0-only WITH Linux-syscall-note */
/*
 * Input event codes
 *
 *    *** IMPORTANT ***
 * This file is not only included from C-code but also from devicetree source
 * files. As such this file MUST only contain comments and defines.
 *
 * Copyright (c) 1999-2002 Vojtech Pavlik
 * Copyright (c) 2015 Hans de Goede <hdegoede@redhat.com>
 *
 * This program is free software; you can redistribute it and/or modify it
 * under the terms of the GNU General Public License version 2 as published by
 * the Free Software Foundation.
 */

/*
* Device properties and quirks
 */
const (
	INPUT_PROP_POINTER        = 0x00 /* needs a pointer */
	INPUT_PROP_DIRECT         = 0x01 /* direct input devices */
	INPUT_PROP_BUTTONPAD      = 0x02 /* has button(s) under pad */
	INPUT_PROP_SEMI_MT        = 0x03 /* touch rectangle only */
	INPUT_PROP_TOPBUTTONPAD   = 0x04 /* softbuttons at top of pad */
	INPUT_PROP_POINTING_STICK = 0x05 /* is a pointing stick */
	INPUT_PROP_ACCELEROMETER  = 0x06 /* has accelerometer */

	INPUT_PROP_MAX = 0x1f
	INPUT_PROP_CNT = (INPUT_PROP_MAX + 1)
)

/*
* Event types
 */
const (
	EV_SYN       EV = 0x00
	EV_KEY       EV = 0x01
	EV_REL       EV = 0x02
	EV_ABS       EV = 0x03
	EV_MSC       EV = 0x04
	EV_SW        EV = 0x05
	EV_LED       EV = 0x11
	EV_SND       EV = 0x12
	EV_REP       EV = 0x14
	EV_FF        EV = 0x15
	EV_PWR       EV = 0x16
	EV_FF_STATUS EV = 0x17
	EV_MAX       EV = 0x1f
	EV_CNT       EV = (EV_MAX + 1)
)

/*
* Synchronization events.
 */
const (
	SYN_REPORT    = 0
	SYN_CONFIG    = 1
	SYN_MT_REPORT = 2
	SYN_DROPPED   = 3
	SYN_MAX       = 0xf
	SYN_CNT       = (SYN_MAX + 1)
)

/*
* Keys and buttons
*
* Most of the keys/buttons are modeled after USB HUT 1.12
* (see http://www.usb.org/developers/hidpage).
* Abbreviations in the comments:
* AC - Application Control
* AL - Application Launch Button
* SC - System Control
 */
const (
	KEY_RESERVED   KEY_CODE = 0
	KEY_ESC        KEY_CODE = 1
	KEY_1          KEY_CODE = 2
	KEY_2          KEY_CODE = 3
	KEY_3          KEY_CODE = 4
	KEY_4          KEY_CODE = 5
	KEY_5          KEY_CODE = 6
	KEY_6          KEY_CODE = 7
	KEY_7          KEY_CODE = 8
	KEY_8          KEY_CODE = 9
	KEY_9          KEY_CODE = 10
	KEY_0          KEY_CODE = 11
	KEY_MINUS      KEY_CODE = 12
	KEY_EQUAL      KEY_CODE = 13
	KEY_BACKSPACE  KEY_CODE = 14
	KEY_TAB        KEY_CODE = 15
	KEY_Q          KEY_CODE = 16
	KEY_W          KEY_CODE = 17
	KEY_E          KEY_CODE = 18
	KEY_R          KEY_CODE = 19
	KEY_T          KEY_CODE = 20
	KEY_Y          KEY_CODE = 21
	KEY_U          KEY_CODE = 22
	KEY_I          KEY_CODE = 23
	KEY_O          KEY_CODE = 24
	KEY_P          KEY_CODE = 25
	KEY_LEFTBRACE  KEY_CODE = 26
	KEY_RIGHTBRACE KEY_CODE = 27
	KEY_ENTER      KEY_CODE = 28
	KEY_LEFTCTRL   KEY_CODE = 29
	KEY_A          KEY_CODE = 30
	KEY_S          KEY_CODE = 31
	KEY_D          KEY_CODE = 32
	KEY_F          KEY_CODE = 33
	KEY_G          KEY_CODE = 34
	KEY_H          KEY_CODE = 35
	KEY_J          KEY_CODE = 36
	KEY_K          KEY_CODE = 37
	KEY_L          KEY_CODE = 38
	KEY_SEMICOLON  KEY_CODE = 39
	KEY_APOSTROPHE KEY_CODE = 40
	KEY_GRAVE      KEY_CODE = 41
	KEY_LEFTSHIFT  KEY_CODE = 42
	KEY_BACKSLASH  KEY_CODE = 43
	KEY_Z          KEY_CODE = 44
	KEY_X          KEY_CODE = 45
	KEY_C          KEY_CODE = 46
	KEY_V          KEY_CODE = 47
	KEY_B          KEY_CODE = 48
	KEY_N          KEY_CODE = 49
	KEY_M          KEY_CODE = 50
	KEY_COMMA      KEY_CODE = 51
	KEY_DOT        KEY_CODE = 52
	KEY_SLASH      KEY_CODE = 53
	KEY_RIGHTSHIFT KEY_CODE = 54
	KEY_KPASTERISK KEY_CODE = 55
	KEY_LEFTALT    KEY_CODE = 56
	KEY_SPACE      KEY_CODE = 57
	KEY_CAPSLOCK   KEY_CODE = 58
	KEY_F1         KEY_CODE = 59
	KEY_F2         KEY_CODE = 60
	KEY_F3         KEY_CODE = 61
	KEY_F4         KEY_CODE = 62
	KEY_F5         KEY_CODE = 63
	KEY_F6         KEY_CODE = 64
	KEY_F7         KEY_CODE = 65
	KEY_F8         KEY_CODE = 66
	KEY_F9         KEY_CODE = 67
	KEY_F10        KEY_CODE = 68
	KEY_NUMLOCK    KEY_CODE = 69
	KEY_SCROLLLOCK KEY_CODE = 70
	KEY_KP7        KEY_CODE = 71
	KEY_KP8        KEY_CODE = 72
	KEY_KP9        KEY_CODE = 73
	KEY_KPMINUS    KEY_CODE = 74
	KEY_KP4        KEY_CODE = 75
	KEY_KP5        KEY_CODE = 76
	KEY_KP6        KEY_CODE = 77
	KEY_KPPLUS     KEY_CODE = 78
	KEY_KP1        KEY_CODE = 79
	KEY_KP2        KEY_CODE = 80
	KEY_KP3        KEY_CODE = 81
	KEY_KP0        KEY_CODE = 82
	KEY_KPDOT      KEY_CODE = 83

	KEY_ZENKAKUHANKAKU   KEY_CODE = 85
	KEY_102ND            KEY_CODE = 86
	KEY_F11              KEY_CODE = 87
	KEY_F12              KEY_CODE = 88
	KEY_RO               KEY_CODE = 89
	KEY_KATAKANA         KEY_CODE = 90
	KEY_HIRAGANA         KEY_CODE = 91
	KEY_HENKAN           KEY_CODE = 92
	KEY_KATAKANAHIRAGANA KEY_CODE = 93
	KEY_MUHENKAN         KEY_CODE = 94
	KEY_KPJPCOMMA        KEY_CODE = 95
	KEY_KPENTER          KEY_CODE = 96
	KEY_RIGHTCTRL        KEY_CODE = 97
	KEY_KPSLASH          KEY_CODE = 98
	KEY_SYSRQ            KEY_CODE = 99
	KEY_RIGHTALT         KEY_CODE = 100
	KEY_LINEFEED         KEY_CODE = 101
	KEY_HOME             KEY_CODE = 102
	KEY_UP               KEY_CODE = 103
	KEY_PAGEUP           KEY_CODE = 104
	KEY_LEFT             KEY_CODE = 105
	KEY_RIGHT            KEY_CODE = 106
	KEY_END              KEY_CODE = 107
	KEY_DOWN             KEY_CODE = 108
	KEY_PAGEDOWN         KEY_CODE = 109
	KEY_INSERT           KEY_CODE = 110
	KEY_DELETE           KEY_CODE = 111
	KEY_MACRO            KEY_CODE = 112
	KEY_MUTE             KEY_CODE = 113
	KEY_VOLUMEDOWN       KEY_CODE = 114
	KEY_VOLUMEUP         KEY_CODE = 115
	KEY_POWER            KEY_CODE = 116 /* SC System Power Down */
	KEY_KPEQUAL          KEY_CODE = 117
	KEY_KPPLUSMINUS      KEY_CODE = 118
	KEY_PAUSE            KEY_CODE = 119
	KEY_SCALE            KEY_CODE = 120 /* AL Compiz Scale (Expose) */

	KEY_KPCOMMA   KEY_CODE = 121
	KEY_HANGEUL   KEY_CODE = 122
	KEY_HANGUEL   KEY_CODE = KEY_HANGEUL
	KEY_HANJA     KEY_CODE = 123
	KEY_YEN       KEY_CODE = 124
	KEY_LEFTMETA  KEY_CODE = 125
	KEY_RIGHTMETA KEY_CODE = 126
	KEY_COMPOSE   KEY_CODE = 127

	KEY_STOP           KEY_CODE = 128 /* AC Stop */
	KEY_AGAIN          KEY_CODE = 129
	KEY_PROPS          KEY_CODE = 130 /* AC Properties */
	KEY_UNDO           KEY_CODE = 131 /* AC Undo */
	KEY_FRONT          KEY_CODE = 132
	KEY_COPY           KEY_CODE = 133 /* AC Copy */
	KEY_OPEN           KEY_CODE = 134 /* AC Open */
	KEY_PASTE          KEY_CODE = 135 /* AC Paste */
	KEY_FIND           KEY_CODE = 136 /* AC Search */
	KEY_CUT            KEY_CODE = 137 /* AC Cut */
	KEY_HELP           KEY_CODE = 138 /* AL Integrated Help Center */
	KEY_MENU           KEY_CODE = 139 /* Menu (show menu) */
	KEY_CALC           KEY_CODE = 140 /* AL Calculator */
	KEY_SETUP          KEY_CODE = 141
	KEY_SLEEP          KEY_CODE = 142 /* SC System Sleep */
	KEY_WAKEUP         KEY_CODE = 143 /* System Wake Up */
	KEY_FILE           KEY_CODE = 144 /* AL Local Machine Browser */
	KEY_SENDFILE       KEY_CODE = 145
	KEY_DELETEFILE     KEY_CODE = 146
	KEY_XFER           KEY_CODE = 147
	KEY_PROG1          KEY_CODE = 148
	KEY_PROG2          KEY_CODE = 149
	KEY_WWW            KEY_CODE = 150 /* AL Internet Browser */
	KEY_MSDOS          KEY_CODE = 151
	KEY_COFFEE         KEY_CODE = 152 /* AL Terminal Lock/Screensaver */
	KEY_SCREENLOCK     KEY_CODE = KEY_COFFEE
	KEY_ROTATE_DISPLAY KEY_CODE = 153 /* Display orientation for e.g. tablets */
	KEY_DIRECTION      KEY_CODE = KEY_ROTATE_DISPLAY
	KEY_CYCLEWINDOWS   KEY_CODE = 154
	KEY_MAIL           KEY_CODE = 155
	KEY_BOOKMARKS      KEY_CODE = 156 /* AC Bookmarks */
	KEY_COMPUTER       KEY_CODE = 157
	KEY_BACK           KEY_CODE = 158 /* AC Back */
	KEY_FORWARD        KEY_CODE = 159 /* AC Forward */
	KEY_CLOSECD        KEY_CODE = 160
	KEY_EJECTCD        KEY_CODE = 161
	KEY_EJECTCLOSECD   KEY_CODE = 162
	KEY_NEXTSONG       KEY_CODE = 163
	KEY_PLAYPAUSE      KEY_CODE = 164
	KEY_PREVIOUSSONG   KEY_CODE = 165
	KEY_STOPCD         KEY_CODE = 166
	KEY_RECORD         KEY_CODE = 167
	KEY_REWIND         KEY_CODE = 168
	KEY_PHONE          KEY_CODE = 169 /* Media Select Telephone */
	KEY_ISO            KEY_CODE = 170
	KEY_CONFIG         KEY_CODE = 171 /* AL Consumer Control Configuration */
	KEY_HOMEPAGE       KEY_CODE = 172 /* AC Home */
	KEY_REFRESH        KEY_CODE = 173 /* AC Refresh */
	KEY_EXIT           KEY_CODE = 174 /* AC Exit */
	KEY_MOVE           KEY_CODE = 175
	KEY_EDIT           KEY_CODE = 176
	KEY_SCROLLUP       KEY_CODE = 177
	KEY_SCROLLDOWN     KEY_CODE = 178
	KEY_KPLEFTPAREN    KEY_CODE = 179
	KEY_KPRIGHTPAREN   KEY_CODE = 180
	KEY_NEW            KEY_CODE = 181 /* AC New */
	KEY_REDO           KEY_CODE = 182 /* AC Redo/Repeat */

	KEY_F13 KEY_CODE = 183
	KEY_F14 KEY_CODE = 184
	KEY_F15 KEY_CODE = 185
	KEY_F16 KEY_CODE = 186
	KEY_F17 KEY_CODE = 187
	KEY_F18 KEY_CODE = 188
	KEY_F19 KEY_CODE = 189
	KEY_F20 KEY_CODE = 190
	KEY_F21 KEY_CODE = 191
	KEY_F22 KEY_CODE = 192
	KEY_F23 KEY_CODE = 193
	KEY_F24 KEY_CODE = 194

	KEY_PLAYCD         KEY_CODE = 200
	KEY_PAUSECD        KEY_CODE = 201
	KEY_PROG3          KEY_CODE = 202
	KEY_PROG4          KEY_CODE = 203
	KEY_DASHBOARD      KEY_CODE = 204 /* AL Dashboard */
	KEY_SUSPEND        KEY_CODE = 205
	KEY_CLOSE          KEY_CODE = 206 /* AC Close */
	KEY_PLAY           KEY_CODE = 207
	KEY_FASTFORWARD    KEY_CODE = 208
	KEY_BASSBOOST      KEY_CODE = 209
	KEY_PRINT          KEY_CODE = 210 /* AC Print */
	KEY_HP             KEY_CODE = 211
	KEY_CAMERA         KEY_CODE = 212
	KEY_SOUND          KEY_CODE = 213
	KEY_QUESTION       KEY_CODE = 214
	KEY_EMAIL          KEY_CODE = 215
	KEY_CHAT           KEY_CODE = 216
	KEY_SEARCH         KEY_CODE = 217
	KEY_CONNECT        KEY_CODE = 218
	KEY_FINANCE        KEY_CODE = 219 /* AL Checkbook/Finance */
	KEY_SPORT          KEY_CODE = 220
	KEY_SHOP           KEY_CODE = 221
	KEY_ALTERASE       KEY_CODE = 222
	KEY_CANCEL         KEY_CODE = 223 /* AC Cancel */
	KEY_BRIGHTNESSDOWN KEY_CODE = 224
	KEY_BRIGHTNESSUP   KEY_CODE = 225
	KEY_MEDIA          KEY_CODE = 226

	KEY_SWITCHVIDEOMODE KEY_CODE = 227 /* Cycle between available video
	outputs (Monitor/LCD/TV-out/etc) */
	KEY_KBDILLUMTOGGLE KEY_CODE = 228
	KEY_KBDILLUMDOWN   KEY_CODE = 229
	KEY_KBDILLUMUP     KEY_CODE = 230

	KEY_SEND        KEY_CODE = 231 /* AC Send */
	KEY_REPLY       KEY_CODE = 232 /* AC Reply */
	KEY_FORWARDMAIL KEY_CODE = 233 /* AC Forward Msg */
	KEY_SAVE        KEY_CODE = 234 /* AC Save */
	KEY_DOCUMENTS   KEY_CODE = 235

	KEY_BATTERY KEY_CODE = 236

	KEY_BLUETOOTH KEY_CODE = 237
	KEY_WLAN      KEY_CODE = 238
	KEY_UWB       KEY_CODE = 239

	KEY_UNKNOWN KEY_CODE = 240

	KEY_VIDEO_NEXT       KEY_CODE = 241 /* drive next video source */
	KEY_VIDEO_PREV       KEY_CODE = 242 /* drive previous video source */
	KEY_BRIGHTNESS_CYCLE KEY_CODE = 243 /* brightness up, after max is min */
	KEY_BRIGHTNESS_AUTO  KEY_CODE = 244 /* Set Auto Brightness: manual
	brightness control is off,
	rely on ambient */
	KEY_BRIGHTNESS_ZERO KEY_CODE = KEY_BRIGHTNESS_AUTO
	KEY_DISPLAY_OFF     KEY_CODE = 245 /* display device to off state */

	KEY_WWAN   KEY_CODE = 246 /* Wireless WAN (LTE, UMTS, GSM, etc.) */
	KEY_WIMAX  KEY_CODE = KEY_WWAN
	KEY_RFKILL KEY_CODE = 247 /* Key that controls all radios */

	KEY_MICMUTE KEY_CODE = 248 /* Mute / unmute the microphone */

	/* Code 255 is reserved for special needs of AT keyboard driver */

	BTN_MISC KEY_CODE = 0x100
	BTN_0    KEY_CODE = 0x100
	BTN_1    KEY_CODE = 0x101
	BTN_2    KEY_CODE = 0x102
	BTN_3    KEY_CODE = 0x103
	BTN_4    KEY_CODE = 0x104
	BTN_5    KEY_CODE = 0x105
	BTN_6    KEY_CODE = 0x106
	BTN_7    KEY_CODE = 0x107
	BTN_8    KEY_CODE = 0x108
	BTN_9    KEY_CODE = 0x109

	BTN_MOUSE   KEY_CODE = 0x110
	BTN_LEFT    KEY_CODE = 0x110
	BTN_RIGHT   KEY_CODE = 0x111
	BTN_MIDDLE  KEY_CODE = 0x112
	BTN_SIDE    KEY_CODE = 0x113
	BTN_EXTRA   KEY_CODE = 0x114
	BTN_FORWARD KEY_CODE = 0x115
	BTN_BACK    KEY_CODE = 0x116
	BTN_TASK    KEY_CODE = 0x117

	BTN_JOYSTICK KEY_CODE = 0x120
	BTN_TRIGGER  KEY_CODE = 0x120
	BTN_THUMB    KEY_CODE = 0x121
	BTN_THUMB2   KEY_CODE = 0x122
	BTN_TOP      KEY_CODE = 0x123
	BTN_TOP2     KEY_CODE = 0x124
	BTN_PINKIE   KEY_CODE = 0x125
	BTN_BASE     KEY_CODE = 0x126
	BTN_BASE2    KEY_CODE = 0x127
	BTN_BASE3    KEY_CODE = 0x128
	BTN_BASE4    KEY_CODE = 0x129
	BTN_BASE5    KEY_CODE = 0x12a
	BTN_BASE6    KEY_CODE = 0x12b
	BTN_DEAD     KEY_CODE = 0x12f

	BTN_GAMEPAD KEY_CODE = 0x130
	BTN_SOUTH   KEY_CODE = 0x130
	BTN_A       KEY_CODE = BTN_SOUTH
	BTN_EAST    KEY_CODE = 0x131
	BTN_B       KEY_CODE = BTN_EAST
	BTN_C       KEY_CODE = 0x132
	BTN_NORTH   KEY_CODE = 0x133
	BTN_X       KEY_CODE = BTN_NORTH
	BTN_WEST    KEY_CODE = 0x134
	BTN_Y       KEY_CODE = BTN_WEST
	BTN_Z       KEY_CODE = 0x135
	BTN_TL      KEY_CODE = 0x136
	BTN_TR      KEY_CODE = 0x137
	BTN_TL2     KEY_CODE = 0x138
	BTN_TR2     KEY_CODE = 0x139
	BTN_SELECT  KEY_CODE = 0x13a
	BTN_START   KEY_CODE = 0x13b
	BTN_MODE    KEY_CODE = 0x13c
	BTN_THUMBL  KEY_CODE = 0x13d
	BTN_THUMBR  KEY_CODE = 0x13e

	BTN_DIGI           KEY_CODE = 0x140
	BTN_TOOL_PEN       KEY_CODE = 0x140
	BTN_TOOL_RUBBER    KEY_CODE = 0x141
	BTN_TOOL_BRUSH     KEY_CODE = 0x142
	BTN_TOOL_PENCIL    KEY_CODE = 0x143
	BTN_TOOL_AIRBRUSH  KEY_CODE = 0x144
	BTN_TOOL_FINGER    KEY_CODE = 0x145
	BTN_TOOL_MOUSE     KEY_CODE = 0x146
	BTN_TOOL_LENS      KEY_CODE = 0x147
	BTN_TOOL_QUINTTAP           = 0x148 /* Five fingers on  KEY_CODEtrackpad */
	BTN_STYLUS3        KEY_CODE = 0x149
	BTN_TOUCH          KEY_CODE = 0x14a
	BTN_STYLUS         KEY_CODE = 0x14b
	BTN_STYLUS2        KEY_CODE = 0x14c
	BTN_TOOL_DOUBLETAP KEY_CODE = 0x14d
	BTN_TOOL_TRIPLETAP KEY_CODE = 0x14e
	BTN_TOOL_QUADTAP            = 0x14f /* Four fingers on  KEY_CODEtrackpad */

	BTN_WHEEL     KEY_CODE = 0x150
	BTN_GEAR_DOWN KEY_CODE = 0x150
	BTN_GEAR_UP   KEY_CODE = 0x151

	KEY_OK                KEY_CODE = 0x160
	KEY_SELECT            KEY_CODE = 0x161
	KEY_GOTO              KEY_CODE = 0x162
	KEY_CLEAR             KEY_CODE = 0x163
	KEY_POWER2            KEY_CODE = 0x164
	KEY_OPTION            KEY_CODE = 0x165
	KEY_INFO              KEY_CODE = 0x166 /* AL OEM Features/Tips/Tutorial */
	KEY_TIME              KEY_CODE = 0x167
	KEY_VENDOR            KEY_CODE = 0x168
	KEY_ARCHIVE           KEY_CODE = 0x169
	KEY_PROGRAM           KEY_CODE = 0x16a /* Media Select Program Guide */
	KEY_CHANNEL           KEY_CODE = 0x16b
	KEY_FAVORITES         KEY_CODE = 0x16c
	KEY_EPG               KEY_CODE = 0x16d
	KEY_PVR               KEY_CODE = 0x16e /* Media Select Home */
	KEY_MHP               KEY_CODE = 0x16f
	KEY_LANGUAGE          KEY_CODE = 0x170
	KEY_TITLE             KEY_CODE = 0x171
	KEY_SUBTITLE          KEY_CODE = 0x172
	KEY_ANGLE             KEY_CODE = 0x173
	KEY_FULL_SCREEN       KEY_CODE = 0x174 /* AC View Toggle */
	KEY_ZOOM              KEY_CODE = KEY_FULL_SCREEN
	KEY_MODE              KEY_CODE = 0x175
	KEY_KEYBOARD          KEY_CODE = 0x176
	KEY_ASPECT_RATIO      KEY_CODE = 0x177 /* HUTRR37: Aspect */
	KEY_SCREEN            KEY_CODE = KEY_ASPECT_RATIO
	KEY_PC                KEY_CODE = 0x178 /* Media Select Computer */
	KEY_TV                KEY_CODE = 0x179 /* Media Select TV */
	KEY_TV2               KEY_CODE = 0x17a /* Media Select Cable */
	KEY_VCR               KEY_CODE = 0x17b /* Media Select VCR */
	KEY_VCR2              KEY_CODE = 0x17c /* VCR Plus */
	KEY_SAT               KEY_CODE = 0x17d /* Media Select Satellite */
	KEY_SAT2              KEY_CODE = 0x17e
	KEY_CD                KEY_CODE = 0x17f /* Media Select CD */
	KEY_TAPE              KEY_CODE = 0x180 /* Media Select Tape */
	KEY_RADIO             KEY_CODE = 0x181
	KEY_TUNER             KEY_CODE = 0x182 /* Media Select Tuner */
	KEY_PLAYER            KEY_CODE = 0x183
	KEY_TEXT              KEY_CODE = 0x184
	KEY_DVD               KEY_CODE = 0x185 /* Media Select DVD */
	KEY_AUX               KEY_CODE = 0x186
	KEY_MP3               KEY_CODE = 0x187
	KEY_AUDIO             KEY_CODE = 0x188 /* AL Audio Browser */
	KEY_VIDEO             KEY_CODE = 0x189 /* AL Movie Browser */
	KEY_DIRECTORY         KEY_CODE = 0x18a
	KEY_LIST              KEY_CODE = 0x18b
	KEY_MEMO              KEY_CODE = 0x18c /* Media Select Messages */
	KEY_CALENDAR          KEY_CODE = 0x18d
	KEY_RED               KEY_CODE = 0x18e
	KEY_GREEN             KEY_CODE = 0x18f
	KEY_YELLOW            KEY_CODE = 0x190
	KEY_BLUE              KEY_CODE = 0x191
	KEY_CHANNELUP         KEY_CODE = 0x192 /* Channel Increment */
	KEY_CHANNELDOWN       KEY_CODE = 0x193 /* Channel Decrement */
	KEY_FIRST             KEY_CODE = 0x194
	KEY_LAST              KEY_CODE = 0x195 /* Recall Last */
	KEY_AB                KEY_CODE = 0x196
	KEY_NEXT              KEY_CODE = 0x197
	KEY_RESTART           KEY_CODE = 0x198
	KEY_SLOW              KEY_CODE = 0x199
	KEY_SHUFFLE           KEY_CODE = 0x19a
	KEY_BREAK             KEY_CODE = 0x19b
	KEY_PREVIOUS          KEY_CODE = 0x19c
	KEY_DIGITS            KEY_CODE = 0x19d
	KEY_TEEN              KEY_CODE = 0x19e
	KEY_TWEN              KEY_CODE = 0x19f
	KEY_VIDEOPHONE        KEY_CODE = 0x1a0 /* Media Select Video Phone */
	KEY_GAMES             KEY_CODE = 0x1a1 /* Media Select Games */
	KEY_ZOOMIN            KEY_CODE = 0x1a2 /* AC Zoom In */
	KEY_ZOOMOUT           KEY_CODE = 0x1a3 /* AC Zoom Out */
	KEY_ZOOMRESET         KEY_CODE = 0x1a4 /* AC Zoom */
	KEY_WORDPROCESSOR     KEY_CODE = 0x1a5 /* AL Word Processor */
	KEY_EDITOR            KEY_CODE = 0x1a6 /* AL Text Editor */
	KEY_SPREADSHEET       KEY_CODE = 0x1a7 /* AL Spreadsheet */
	KEY_GRAPHICSEDITOR    KEY_CODE = 0x1a8 /* AL Graphics Editor */
	KEY_PRESENTATION      KEY_CODE = 0x1a9 /* AL Presentation App */
	KEY_DATABASE          KEY_CODE = 0x1aa /* AL Database App */
	KEY_NEWS              KEY_CODE = 0x1ab /* AL Newsreader */
	KEY_VOICEMAIL         KEY_CODE = 0x1ac /* AL Voicemail */
	KEY_ADDRESSBOOK       KEY_CODE = 0x1ad /* AL Contacts/Address Book */
	KEY_MESSENGER         KEY_CODE = 0x1ae /* AL Instant Messaging */
	KEY_DISPLAYTOGGLE     KEY_CODE = 0x1af /* Turn display (LCD) on and off */
	KEY_BRIGHTNESS_TOGGLE KEY_CODE = KEY_DISPLAYTOGGLE
	KEY_SPELLCHECK        KEY_CODE = 0x1b0 /* AL Spell Check */
	KEY_LOGOFF            KEY_CODE = 0x1b1 /* AL Logoff */

	KEY_DOLLAR KEY_CODE = 0x1b2
	KEY_EURO   KEY_CODE = 0x1b3

	KEY_FRAMEBACK           KEY_CODE = 0x1b4 /* Consumer - transport controls */
	KEY_FRAMEFORWARD        KEY_CODE = 0x1b5
	KEY_CONTEXT_MENU        KEY_CODE = 0x1b6 /* GenDesc - system context menu */
	KEY_MEDIA_REPEAT        KEY_CODE = 0x1b7 /* Consumer - transport control */
	KEY_10CHANNELSUP        KEY_CODE = 0x1b8 /* 10 channels up (10+) */
	KEY_10CHANNELSDOWN      KEY_CODE = 0x1b9 /* 10 channels down (10-) */
	KEY_IMAGES              KEY_CODE = 0x1ba /* AL Image Browser */
	KEY_NOTIFICATION_CENTER KEY_CODE = 0x1bc /* Show/hide the notification center */
	KEY_PICKUP_PHONE        KEY_CODE = 0x1bd /* Answer incoming call */
	KEY_HANGUP_PHONE        KEY_CODE = 0x1be /* Decline incoming call */

	KEY_DEL_EOL  KEY_CODE = 0x1c0
	KEY_DEL_EOS  KEY_CODE = 0x1c1
	KEY_INS_LINE KEY_CODE = 0x1c2
	KEY_DEL_LINE KEY_CODE = 0x1c3

	KEY_FN             KEY_CODE = 0x1d0
	KEY_FN_ESC         KEY_CODE = 0x1d1
	KEY_FN_F1          KEY_CODE = 0x1d2
	KEY_FN_F2          KEY_CODE = 0x1d3
	KEY_FN_F3          KEY_CODE = 0x1d4
	KEY_FN_F4          KEY_CODE = 0x1d5
	KEY_FN_F5          KEY_CODE = 0x1d6
	KEY_FN_F6          KEY_CODE = 0x1d7
	KEY_FN_F7          KEY_CODE = 0x1d8
	KEY_FN_F8          KEY_CODE = 0x1d9
	KEY_FN_F9          KEY_CODE = 0x1da
	KEY_FN_F10         KEY_CODE = 0x1db
	KEY_FN_F11         KEY_CODE = 0x1dc
	KEY_FN_F12         KEY_CODE = 0x1dd
	KEY_FN_1           KEY_CODE = 0x1de
	KEY_FN_2           KEY_CODE = 0x1df
	KEY_FN_D           KEY_CODE = 0x1e0
	KEY_FN_E           KEY_CODE = 0x1e1
	KEY_FN_F           KEY_CODE = 0x1e2
	KEY_FN_S           KEY_CODE = 0x1e3
	KEY_FN_B           KEY_CODE = 0x1e4
	KEY_FN_RIGHT_SHIFT KEY_CODE = 0x1e5

	KEY_BRL_DOT1  KEY_CODE = 0x1f1
	KEY_BRL_DOT2  KEY_CODE = 0x1f2
	KEY_BRL_DOT3  KEY_CODE = 0x1f3
	KEY_BRL_DOT4  KEY_CODE = 0x1f4
	KEY_BRL_DOT5  KEY_CODE = 0x1f5
	KEY_BRL_DOT6  KEY_CODE = 0x1f6
	KEY_BRL_DOT7  KEY_CODE = 0x1f7
	KEY_BRL_DOT8  KEY_CODE = 0x1f8
	KEY_BRL_DOT9  KEY_CODE = 0x1f9
	KEY_BRL_DOT10 KEY_CODE = 0x1fa

	KEY_NUMERIC_0     KEY_CODE = 0x200 /* used by phones, remote controls, */
	KEY_NUMERIC_1     KEY_CODE = 0x201 /* and other keypads */
	KEY_NUMERIC_2     KEY_CODE = 0x202
	KEY_NUMERIC_3     KEY_CODE = 0x203
	KEY_NUMERIC_4     KEY_CODE = 0x204
	KEY_NUMERIC_5     KEY_CODE = 0x205
	KEY_NUMERIC_6     KEY_CODE = 0x206
	KEY_NUMERIC_7     KEY_CODE = 0x207
	KEY_NUMERIC_8     KEY_CODE = 0x208
	KEY_NUMERIC_9     KEY_CODE = 0x209
	KEY_NUMERIC_STAR  KEY_CODE = 0x20a
	KEY_NUMERIC_POUND KEY_CODE = 0x20b
	KEY_NUMERIC_A     KEY_CODE = 0x20c /* Phone key A - HUT Telephony 0xb9 */
	KEY_NUMERIC_B     KEY_CODE = 0x20d
	KEY_NUMERIC_C     KEY_CODE = 0x20e
	KEY_NUMERIC_D     KEY_CODE = 0x20f

	KEY_CAMERA_FOCUS KEY_CODE = 0x210
	KEY_WPS_BUTTON   KEY_CODE = 0x211 /* WiFi Protected Setup key */

	KEY_TOUCHPAD_TOGGLE KEY_CODE = 0x212 /* Request switch touchpad on or off */
	KEY_TOUCHPAD_ON     KEY_CODE = 0x213
	KEY_TOUCHPAD_OFF    KEY_CODE = 0x214

	KEY_CAMERA_ZOOMIN  KEY_CODE = 0x215
	KEY_CAMERA_ZOOMOUT KEY_CODE = 0x216
	KEY_CAMERA_UP      KEY_CODE = 0x217
	KEY_CAMERA_DOWN    KEY_CODE = 0x218
	KEY_CAMERA_LEFT    KEY_CODE = 0x219
	KEY_CAMERA_RIGHT   KEY_CODE = 0x21a

	KEY_ATTENDANT_ON     KEY_CODE = 0x21b
	KEY_ATTENDANT_OFF    KEY_CODE = 0x21c
	KEY_ATTENDANT_TOGGLE KEY_CODE = 0x21d /* Attendant call on or off */
	KEY_LIGHTS_TOGGLE    KEY_CODE = 0x21e /* Reading light on or off */

	BTN_DPAD_UP    KEY_CODE = 0x220
	BTN_DPAD_DOWN  KEY_CODE = 0x221
	BTN_DPAD_LEFT  KEY_CODE = 0x222
	BTN_DPAD_RIGHT KEY_CODE = 0x223

	KEY_ALS_TOGGLE         KEY_CODE = 0x230 /* Ambient light sensor */
	KEY_ROTATE_LOCK_TOGGLE KEY_CODE = 0x231 /* Display rotation lock */

	KEY_BUTTONCONFIG    KEY_CODE = 0x240 /* AL Button Configuration */
	KEY_TASKMANAGER     KEY_CODE = 0x241 /* AL Task/Project Manager */
	KEY_JOURNAL         KEY_CODE = 0x242 /* AL Log/Journal/Timecard */
	KEY_CONTROLPANEL    KEY_CODE = 0x243 /* AL Control Panel */
	KEY_APPSELECT       KEY_CODE = 0x244 /* AL Select Task/Application */
	KEY_SCREENSAVER     KEY_CODE = 0x245 /* AL Screen Saver */
	KEY_VOICECOMMAND    KEY_CODE = 0x246 /* Listening Voice Command */
	KEY_ASSISTANT       KEY_CODE = 0x247 /* AL Context-aware desktop assistant */
	KEY_KBD_LAYOUT_NEXT KEY_CODE = 0x248 /* AC Next Keyboard Layout Select */
	KEY_EMOJI_PICKER    KEY_CODE = 0x249 /* Show/hide emoji picker (HUTRR101) */

	KEY_BRIGHTNESS_MIN KEY_CODE = 0x250 /* Set Brightness to Minimum */
	KEY_BRIGHTNESS_MAX KEY_CODE = 0x251 /* Set Brightness to Maximum */

	KEY_KBDINPUTASSIST_PREV      KEY_CODE = 0x260
	KEY_KBDINPUTASSIST_NEXT      KEY_CODE = 0x261
	KEY_KBDINPUTASSIST_PREVGROUP KEY_CODE = 0x262
	KEY_KBDINPUTASSIST_NEXTGROUP KEY_CODE = 0x263
	KEY_KBDINPUTASSIST_ACCEPT    KEY_CODE = 0x264
	KEY_KBDINPUTASSIST_CANCEL    KEY_CODE = 0x265

	/* Diagonal movement keys */
	KEY_RIGHT_UP   KEY_CODE = 0x266
	KEY_RIGHT_DOWN KEY_CODE = 0x267
	KEY_LEFT_UP    KEY_CODE = 0x268
	KEY_LEFT_DOWN  KEY_CODE = 0x269

	KEY_ROOT_MENU KEY_CODE = 0x26a /* Show Device's Root Menu */
	/* Show Top Menu of the Media (e.g. DVD) */
	KEY_MEDIA_TOP_MENU KEY_CODE = 0x26b
	KEY_NUMERIC_11     KEY_CODE = 0x26c
	KEY_NUMERIC_12     KEY_CODE = 0x26d
	/*
	* Toggle Audio Description: refers to an audio service that helps blind and
	* visually impaired consumers understand the action in a program. Note: in
	* some countries this is referred to as "Video Description".
	 */
	KEY_AUDIO_DESC    KEY_CODE = 0x26e
	KEY_3D_MODE       KEY_CODE = 0x26f
	KEY_NEXT_FAVORITE KEY_CODE = 0x270
	KEY_STOP_RECORD   KEY_CODE = 0x271
	KEY_PAUSE_RECORD  KEY_CODE = 0x272
	KEY_VOD           KEY_CODE = 0x273 /* Video on Demand */
	KEY_UNMUTE        KEY_CODE = 0x274
	KEY_FASTREVERSE   KEY_CODE = 0x275
	KEY_SLOWREVERSE   KEY_CODE = 0x276
	/*
	* Control a data application associated with the currently viewed channel,
	* e.g. teletext or data broadcast application (MHEG, MHP, HbbTV, etc.)
	 */
	KEY_DATA              KEY_CODE = 0x277
	KEY_ONSCREEN_KEYBOARD KEY_CODE = 0x278
	/* Electronic privacy screen control */
	KEY_PRIVACY_SCREEN_TOGGLE KEY_CODE = 0x279

	/* Select an area of screen to be copied */
	KEY_SELECTIVE_SCREENSHOT KEY_CODE = 0x27a

	/*
	* Some keyboards have keys which do not have a defined meaning, these keys
	* are intended to be programmed / bound to macros by the user. For most
	* keyboards with these macro-keys the key-sequence to inject, or action to
	* take, is all handled by software on the host side. So from the kernel's
	* point of view these are just normal keys.
	*
	* The KEY_MACRO KEY_CODE# codes below are intended for such keys, which may be labeled
	* e.g. G1-G18, or S1 - S30. The KEY_MACRO KEY_CODE# codes MUST NOT be used for keys
	* where the marking on the key does indicate a defined meaning / purpose.
	*
	* The KEY_MACRO KEY_CODE# codes MUST also NOT be used as fallback for when no existing
	* KEY_FOO KEY_CODE define matches the marking / purpose. In this case a new KEY_FOO KEY_CODE
	* define MUST be added.
	 */
	KEY_MACRO1  KEY_CODE = 0x290
	KEY_MACRO2  KEY_CODE = 0x291
	KEY_MACRO3  KEY_CODE = 0x292
	KEY_MACRO4  KEY_CODE = 0x293
	KEY_MACRO5  KEY_CODE = 0x294
	KEY_MACRO6  KEY_CODE = 0x295
	KEY_MACRO7  KEY_CODE = 0x296
	KEY_MACRO8  KEY_CODE = 0x297
	KEY_MACRO9  KEY_CODE = 0x298
	KEY_MACRO10 KEY_CODE = 0x299
	KEY_MACRO11 KEY_CODE = 0x29a
	KEY_MACRO12 KEY_CODE = 0x29b
	KEY_MACRO13 KEY_CODE = 0x29c
	KEY_MACRO14 KEY_CODE = 0x29d
	KEY_MACRO15 KEY_CODE = 0x29e
	KEY_MACRO16 KEY_CODE = 0x29f
	KEY_MACRO17 KEY_CODE = 0x2a0
	KEY_MACRO18 KEY_CODE = 0x2a1
	KEY_MACRO19 KEY_CODE = 0x2a2
	KEY_MACRO20 KEY_CODE = 0x2a3
	KEY_MACRO21 KEY_CODE = 0x2a4
	KEY_MACRO22 KEY_CODE = 0x2a5
	KEY_MACRO23 KEY_CODE = 0x2a6
	KEY_MACRO24 KEY_CODE = 0x2a7
	KEY_MACRO25 KEY_CODE = 0x2a8
	KEY_MACRO26 KEY_CODE = 0x2a9
	KEY_MACRO27 KEY_CODE = 0x2aa
	KEY_MACRO28 KEY_CODE = 0x2ab
	KEY_MACRO29 KEY_CODE = 0x2ac
	KEY_MACRO30 KEY_CODE = 0x2ad

	/*
	* Some keyboards with the macro-keys described above have some extra keys
	* for controlling the host-side software responsible for the macro handling:
	* -A macro recording start/stop key. Note that not all keyboards which emit
	*  KEY_MACRO_RECORD_START KEY_CODE will also emit KEY_MACRO_RECORD_STOP KEY_CODE if
	*  KEY_MACRO_RECORD_STOP KEY_CODE is not advertised, then KEY_MACRO_RECORD_START KEY_CODE
	*  should be interpreted as a recording start/stop toggle;
	* -Keys for switching between different macro (pre)sets, either a key for
	*  cycling through the configured presets or keys to directly select a preset.
	 */
	KEY_MACRO_RECORD_START KEY_CODE = 0x2b0
	KEY_MACRO_RECORD_STOP  KEY_CODE = 0x2b1
	KEY_MACRO_PRESET_CYCLE KEY_CODE = 0x2b2
	KEY_MACRO_PRESET1      KEY_CODE = 0x2b3
	KEY_MACRO_PRESET2      KEY_CODE = 0x2b4
	KEY_MACRO_PRESET3      KEY_CODE = 0x2b5

	/*
	* Some keyboards have a buildin LCD panel where the contents are controlled
	* by the host. Often these have a number of keys directly below the LCD
	* intended for controlling a menu shown on the LCD. These keys often don't
	* have any labeling so we just name them KEY_KBD_LCD_MENU KEY_CODE#
	 */
	KEY_KBD_LCD_MENU1 KEY_CODE = 0x2b8
	KEY_KBD_LCD_MENU2 KEY_CODE = 0x2b9
	KEY_KBD_LCD_MENU3 KEY_CODE = 0x2ba
	KEY_KBD_LCD_MENU4 KEY_CODE = 0x2bb
	KEY_KBD_LCD_MENU5 KEY_CODE = 0x2bc

	BTN_TRIGGER_HAPPY   KEY_CODE = 0x2c0
	BTN_TRIGGER_HAPPY1  KEY_CODE = 0x2c0
	BTN_TRIGGER_HAPPY2  KEY_CODE = 0x2c1
	BTN_TRIGGER_HAPPY3  KEY_CODE = 0x2c2
	BTN_TRIGGER_HAPPY4  KEY_CODE = 0x2c3
	BTN_TRIGGER_HAPPY5  KEY_CODE = 0x2c4
	BTN_TRIGGER_HAPPY6  KEY_CODE = 0x2c5
	BTN_TRIGGER_HAPPY7  KEY_CODE = 0x2c6
	BTN_TRIGGER_HAPPY8  KEY_CODE = 0x2c7
	BTN_TRIGGER_HAPPY9  KEY_CODE = 0x2c8
	BTN_TRIGGER_HAPPY10 KEY_CODE = 0x2c9
	BTN_TRIGGER_HAPPY11 KEY_CODE = 0x2ca
	BTN_TRIGGER_HAPPY12 KEY_CODE = 0x2cb
	BTN_TRIGGER_HAPPY13 KEY_CODE = 0x2cc
	BTN_TRIGGER_HAPPY14 KEY_CODE = 0x2cd
	BTN_TRIGGER_HAPPY15 KEY_CODE = 0x2ce
	BTN_TRIGGER_HAPPY16 KEY_CODE = 0x2cf
	BTN_TRIGGER_HAPPY17 KEY_CODE = 0x2d0
	BTN_TRIGGER_HAPPY18 KEY_CODE = 0x2d1
	BTN_TRIGGER_HAPPY19 KEY_CODE = 0x2d2
	BTN_TRIGGER_HAPPY20 KEY_CODE = 0x2d3
	BTN_TRIGGER_HAPPY21 KEY_CODE = 0x2d4
	BTN_TRIGGER_HAPPY22 KEY_CODE = 0x2d5
	BTN_TRIGGER_HAPPY23 KEY_CODE = 0x2d6
	BTN_TRIGGER_HAPPY24 KEY_CODE = 0x2d7
	BTN_TRIGGER_HAPPY25 KEY_CODE = 0x2d8
	BTN_TRIGGER_HAPPY26 KEY_CODE = 0x2d9
	BTN_TRIGGER_HAPPY27 KEY_CODE = 0x2da
	BTN_TRIGGER_HAPPY28 KEY_CODE = 0x2db
	BTN_TRIGGER_HAPPY29 KEY_CODE = 0x2dc
	BTN_TRIGGER_HAPPY30 KEY_CODE = 0x2dd
	BTN_TRIGGER_HAPPY31 KEY_CODE = 0x2de
	BTN_TRIGGER_HAPPY32 KEY_CODE = 0x2df
	BTN_TRIGGER_HAPPY33 KEY_CODE = 0x2e0
	BTN_TRIGGER_HAPPY34 KEY_CODE = 0x2e1
	BTN_TRIGGER_HAPPY35 KEY_CODE = 0x2e2
	BTN_TRIGGER_HAPPY36 KEY_CODE = 0x2e3
	BTN_TRIGGER_HAPPY37 KEY_CODE = 0x2e4
	BTN_TRIGGER_HAPPY38 KEY_CODE = 0x2e5
	BTN_TRIGGER_HAPPY39 KEY_CODE = 0x2e6
	BTN_TRIGGER_HAPPY40 KEY_CODE = 0x2e7

	/* We avoid low common keys in module aliases so they don't get huge. */
	KEY_MIN_INTERESTING KEY_CODE = KEY_MUTE
	KEY_MAX             KEY_CODE = 0x2ff
	KEY_CNT             KEY_CODE = (KEY_MAX + 1)
)

/*
* Relative axes
 */
const (
	REL_X      REL_CODE = 0x00
	REL_Y      REL_CODE = 0x01
	REL_Z      REL_CODE = 0x02
	REL_RX     REL_CODE = 0x03
	REL_RY     REL_CODE = 0x04
	REL_RZ     REL_CODE = 0x05
	REL_HWHEEL REL_CODE = 0x06
	REL_DIAL   REL_CODE = 0x07
	REL_WHEEL  REL_CODE = 0x08
	REL_MISC   REL_CODE = 0x09
	/*
	* 0x0a is reserved and should not be used in input drivers.
	* It was used by HID as REL_MISC+1 and userspace needs to detect if
	* the next REL_* event is correct or is just REL_MISC + n.
	* We define here REL_RESERVED so userspace can rely on it and detect
	* the situation described above.
	 */
	REL_RESERVED      REL_CODE = 0x0a
	REL_WHEEL_HI_RES  REL_CODE = 0x0b
	REL_HWHEEL_HI_RES REL_CODE = 0x0c
	REL_MAX           REL_CODE = 0x0f
	REL_CNT           REL_CODE = (REL_MAX + 1)
)

/*
* Absolute axes
 */
const (
	ABS_X          = 0x00
	ABS_Y          = 0x01
	ABS_Z          = 0x02
	ABS_RX         = 0x03
	ABS_RY         = 0x04
	ABS_RZ         = 0x05
	ABS_THROTTLE   = 0x06
	ABS_RUDDER     = 0x07
	ABS_WHEEL      = 0x08
	ABS_GAS        = 0x09
	ABS_BRAKE      = 0x0a
	ABS_HAT0X      = 0x10
	ABS_HAT0Y      = 0x11
	ABS_HAT1X      = 0x12
	ABS_HAT1Y      = 0x13
	ABS_HAT2X      = 0x14
	ABS_HAT2Y      = 0x15
	ABS_HAT3X      = 0x16
	ABS_HAT3Y      = 0x17
	ABS_PRESSURE   = 0x18
	ABS_DISTANCE   = 0x19
	ABS_TILT_X     = 0x1a
	ABS_TILT_Y     = 0x1b
	ABS_TOOL_WIDTH = 0x1c

	ABS_VOLUME = 0x20

	ABS_MISC = 0x28

	/*
	* 0x2e is reserved and should not be used in input drivers.
	* It was used by HID as ABS_MISC+6 and userspace needs to detect if
	* the next ABS_* event is correct or is just ABS_MISC + n.
	* We define here ABS_RESERVED so userspace can rely on it and detect
	* the situation described above.
	 */
	ABS_RESERVED = 0x2e

	ABS_MT_SLOT        = 0x2f /* MT slot being modified */
	ABS_MT_TOUCH_MAJOR = 0x30 /* Major axis of touching ellipse */
	ABS_MT_TOUCH_MINOR = 0x31 /* Minor axis (omit if circular) */
	ABS_MT_WIDTH_MAJOR = 0x32 /* Major axis of approaching ellipse */
	ABS_MT_WIDTH_MINOR = 0x33 /* Minor axis (omit if circular) */
	ABS_MT_ORIENTATION = 0x34 /* Ellipse orientation */
	ABS_MT_POSITION_X  = 0x35 /* Center X touch position */
	ABS_MT_POSITION_Y  = 0x36 /* Center Y touch position */
	ABS_MT_TOOL_TYPE   = 0x37 /* Type of touching device */
	ABS_MT_BLOB_ID     = 0x38 /* Group a set of packets as a blob */
	ABS_MT_TRACKING_ID = 0x39 /* Unique ID of initiated contact */
	ABS_MT_PRESSURE    = 0x3a /* Pressure on contact area */
	ABS_MT_DISTANCE    = 0x3b /* Contact hover distance */
	ABS_MT_TOOL_X      = 0x3c /* Center X tool position */
	ABS_MT_TOOL_Y      = 0x3d /* Center Y tool position */

	ABS_MAX = 0x3f
	ABS_CNT = (ABS_MAX + 1)
)

/*
* Switch events
 */
const (
	SW_LID              = 0x00 /* set = lid shut */
	SW_TABLET_MODE      = 0x01 /* set = tablet mode */
	SW_HEADPHONE_INSERT = 0x02 /* set = inserted */
	SW_RFKILL_ALL       = 0x03 /* rfkill master switch, type "any"
	set = radio enabled */
	SW_RADIO                = SW_RFKILL_ALL /* deprecated */
	SW_MICROPHONE_INSERT    = 0x04          /* set = inserted */
	SW_DOCK                 = 0x05          /* set = plugged into dock */
	SW_LINEOUT_INSERT       = 0x06          /* set = inserted */
	SW_JACK_PHYSICAL_INSERT = 0x07          /* set = mechanical switch set */
	SW_VIDEOOUT_INSERT      = 0x08          /* set = inserted */
	SW_CAMERA_LENS_COVER    = 0x09          /* set = lens covered */
	SW_KEYPAD_SLIDE         = 0x0a          /* set = keypad slide out */
	SW_FRONT_PROXIMITY      = 0x0b          /* set = front proximity sensor active */
	SW_ROTATE_LOCK          = 0x0c          /* set = rotate locked/disabled */
	SW_LINEIN_INSERT        = 0x0d          /* set = inserted */
	SW_MUTE_DEVICE          = 0x0e          /* set = device disabled */
	SW_PEN_INSERTED         = 0x0f          /* set = pen inserted */
	SW_MACHINE_COVER        = 0x10          /* set = cover closed */
	SW_MAX                  = 0x10
	SW_CNT                  = (SW_MAX + 1)
)

/*
* Misc events
 */
const (
	MSC_SERIAL    = 0x00
	MSC_PULSELED  = 0x01
	MSC_GESTURE   = 0x02
	MSC_RAW       = 0x03
	MSC_SCAN      = 0x04
	MSC_TIMESTAMP = 0x05
	MSC_MAX       = 0x07
	MSC_CNT       = (MSC_MAX + 1)
)

/*
* LEDs
 */
const (
	LED_NUML     = 0x00
	LED_CAPSL    = 0x01
	LED_SCROLLL  = 0x02
	LED_COMPOSE  = 0x03
	LED_KANA     = 0x04
	LED_SLEEP    = 0x05
	LED_SUSPEND  = 0x06
	LED_MUTE     = 0x07
	LED_MISC     = 0x08
	LED_MAIL     = 0x09
	LED_CHARGING = 0x0a
	LED_MAX      = 0x0f
	LED_CNT      = (LED_MAX + 1)
)

/*
* Autorepeat values
 */
const (
	REP_DELAY  = 0x00
	REP_PERIOD = 0x01
	REP_MAX    = 0x01
	REP_CNT    = (REP_MAX + 1)
)

/*
* Sounds
 */
const (
	SND_CLICK = 0x00
	SND_BELL  = 0x01
	SND_TONE  = 0x02
	SND_MAX   = 0x07
	SND_CNT   = (SND_MAX + 1)
)
