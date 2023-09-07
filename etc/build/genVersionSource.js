#!/usr/bin/env node

const nodePath = require('path')

const {BusinessException, string: { cap }} = require('@almadash/shelf')
const {fsShelf, DirFs} = require('@almadash/shelf-node')

/* @info
  - FROM package.json `version` field
  - GENERATE a golang source file
  - that provides a const var with the version
  - for use within the program
  - baked within the compiled binary
*/
function main() {
  const rootDir = nodePath.resolve(`${__dirname}/../..`)
  const packagePath = `${rootDir}/package.json`

  if(!fsShelf.exists(packagePath)) {
    throw new BusinessException(`package.json NOT found at ${packagePath}`)
  }

  const {version} = require(packagePath)

  const text = cap(`
    package auto

    const Version="${version}"
  `)

  const autoDir = new DirFs(`${rootDir}/src/auto`)
  autoDir.ensureBaseDir()

  autoDir.saveFile('meta.go', text)
}

main()
