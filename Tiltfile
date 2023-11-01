if not os.path.exists('../tilt-support'):
    fail('Please clone https://github.com/operator-framework/tilt-support to ../tilt-support')

load('../tilt-support/Tiltfile', 'deploy_repo')

config.define_string_list('repos', args=True)
cfg = config.parse()
repos = cfg.get('repos', ['operator-controller', 'rukpak', 'catalogd'])

repo = {
    'image': 'quay.io/operator-framework/operator-controller',
    'yaml': 'config/default',
    'binaries': {
        'manager': 'operator-controller-controller-manager',
    },
    'starting_debug_port': 30000,
}

for r in repos:
    if r == 'operator-controller':
        deploy_repo('operator-controller', repo)
    else:
        include('../{}/Tiltfile'.format(r))
