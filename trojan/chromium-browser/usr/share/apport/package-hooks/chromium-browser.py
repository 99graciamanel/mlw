'''chromium apport hook

/usr/share/apport/package-hooks/chromium-browser.py

Copyright (c) 2010, Fabien Tassin <fta@sofaraway.org>
Copyright (c) 2014-2019, Canonical

This program is free software; you can redistribute it and/or modify it
under the terms of the GNU General Public License as published by the
Free Software Foundation; either version 2 of the License, or (at your
option) any later version.  See http://www.gnu.org/copyleft/gpl.html for
the full text of the license.
'''

import hashlib
import json
import os
import os.path
import re
import apport.hookutils


SNAP_USER_DATA = os.path.expanduser('~/snap/chromium/current')


def user_prefs(report, filename):
    with open(filename, 'r') as f:
        prefs = json.load(f)
    entries = []
    if 'extensions' in prefs:
        extensions = prefs['extensions']
        if 'theme' in extensions:
            theme = extensions['theme']
            if 'use_system' in theme:
                entries.append('extensions.theme.use_system = {}'.format(
                    theme['use_system']))
        if 'settings' in extensions:
            settings = extensions['settings']
            if settings:
                entries.append('extensions:')
            for key in settings.keys():
                extension = settings[key]
                entries.append('  {}:'.format(key))
                if 'manifest' in extension:
                    manifest = extension['manifest']
                    for mkey in ('name', 'version', 'key'):
                        entries.append('    {} = {}'.format(mkey,
                                                            manifest[mkey]))
                if 'state' in extension:
                    entries.append('    state = {}'.format(extension['state']))
    report['Snap.ChromiumPrefs'] = '\n'.join(entries)


def run_cmd(cmd):
    return apport.hookutils.command_output(cmd)


def add_info(report, hookui):
    for snap in ['core', 'core18', 'chromium', 'gtk-common-themes']:
        report['Snap.Info.{}'.format(snap)] = \
            run_cmd(['snap', 'info', '--abs-time', snap])

    report['Snap.Connections'] = run_cmd(['snap', 'connections', 'chromium'])

    report['Snap.Changes'] = \
        run_cmd(['snap', 'changes', '--abs-time', 'chromium'])

    report['Snap.ChromiumVersion'] = \
        run_cmd(['snap', 'run', 'chromium', '--version'])
    report['Snap.ChromeDriverVersion'] = \
        run_cmd(['snap', 'run', 'chromium.chromedriver', '--version'])

    libsdir = os.path.join(SNAP_USER_DATA, '.local/lib')
    if os.path.exists(libsdir):
        libs = []
        for lib in os.listdir(libsdir):
            fp = os.path.join(libsdir, lib)
            with open(fp, 'rb') as f:
                sha256sum = hashlib.sha256(f.read()).hexdigest()
                libs.append('{} ({})'.format(lib, sha256sum))
        report['Snap.LocalLibs'] = '\n'.join(libs)

    flashlib = os.path.join(libsdir, 'libpepflashplayer.so')
    if os.path.exists(flashlib):
        with open(flashlib, 'rb') as f:
            r = re.compile(b'\x00\\.swf\\.playerversion\x00(.*?)\x00')
            version = r.findall(f.read())
            if len(version) > 0:
                report['Snap.FlashPlayerVersion'] = version[0].decode()

    user_profile_dir = os.path.join(SNAP_USER_DATA, '.config/chromium/Default')
    user_prefs(report, os.path.join(user_profile_dir, 'Preferences'))

    report['DiskUsage'] = \
        run_cmd(['df', '-Th', '/home', '/run/shm', SNAP_USER_DATA])

    apport.hookutils._add_tag(report, 'snap')

    apport.hookutils.attach_hardware(report)
    apport.hookutils.attach_drm_info(report)
    apport.hookutils.attach_dmesg(report)


def _test():
    report = {}
    add_info(report, None)
    for key in report:
        print('[{}]\n{}\n'.format(key, report[key]))


if __name__ == '__main__':
    _test()
