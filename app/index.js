'use strict';

const path = require('path');
const Generator = require('yeoman-generator');
const mkdir = require('mkdirp');

module.exports = class extends Generator {

    paths() {
        this.destinationRoot(process.env.GOPATH || './');
    }

    prompting() {

        console.log('\n' +
            '+-----------------------------------+\n' +
            '| G O   E C H O   G E N E R A T O R |\n' +
            '+-----------------------------------+\n' +
            '\n');

        let cb = this.async();

        let prompts = [{
            type: 'input',
            name: 'appName',
            message: 'What is the name of your application?',
            default: 'myapp'
        }, {
            type: 'input',
            name: 'repoUrl',
            message: 'What is your URL repository?',
            default: 'github.com/me'
        }, {
            type: 'input',
            name: 'portNum',
            message: 'Enter port number?',
            default: '8001'
        }];

        return this.prompt(prompts).then(props => {
            this.appName = props.appName.replace(/\s+/g, '-').toLowerCase();
            if (props.repoUrl.substr(-1) != '/') props.repoUrl += '/';
            this.repoUrl = props.repoUrl + this.appName;
            this.portNum = props.portNum
            cb()
        });

    }

    writing() {
        console.log('Generating tree folders');
        let pkgDir = this.destinationPath('pkg');
        let srcDir = this.destinationPath(path.join('src/', this.repoUrl));
        let binDir = this.destinationPath('bin');

        mkdir.sync(pkgDir);
        mkdir.sync(srcDir);
        mkdir.sync(binDir);

        this.fs.copy(
            this.templatePath('_gitignore'),
            path.join(srcDir, '.gitignore')
        );
        this.fs.copy(
            this.templatePath('assets'),
            path.join(srcDir, '/assets')
        );
        this.fs.copy(
            this.templatePath('config'),
            path.join(srcDir, '/config')
        );
        this.fs.copy(
            this.templatePath('view'),
            path.join(srcDir, '/view')
        );
        this.fs.copy(
            this.templatePath('helpr'),
            path.join(srcDir, '/helpr')
        );
        this.fs.copy(
            this.templatePath('model'),
            path.join(srcDir, '/model')
        );
        this.fs.copy(
            this.templatePath('controller/_test.go'),
            path.join(srcDir, '/controller/_test.go'),
        );
        this.fs.copy(
            this.templatePath('controller/router.go'),
            path.join(srcDir, '/controller/router.go'),
        );
        this.fs.copy(
            this.templatePath('controller/index_test.go'),
            path.join(srcDir, '/controller/index_test.go'),
        );
        this.fs.copy(
            this.templatePath('controller/index.go'),
            path.join(srcDir, '/controller/index.go'),
        );

        let tmplContext = {
            appName: this.appName,
            repoUrl: this.repoUrl,
            portNum: this.portNum,
        };

        this.fs.copyTpl(
            this.templatePath('go.mod'),
            path.join(srcDir, 'go.mod'),
            tmplContext
        );
        this.fs.copyTpl(
            this.templatePath('_envrc'),
            path.join(srcDir, '.envrc'),
            tmplContext
        );
        this.fs.copyTpl(
            this.templatePath('main.go'),
            path.join(srcDir, 'main.go'),
            tmplContext
        );
        this.fs.copyTpl(
            this.templatePath('init.go'),
            path.join(srcDir, 'init.go'),
            tmplContext
        );
        this.fs.copyTpl(
            this.templatePath('view/_template.html'),
            path.join(srcDir, '/view/_template.html'),
            tmplContext
        );
    }
};
