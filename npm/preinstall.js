const fs = require("fs");

const supported = {
  arch: {
    ia32: "386",
    x64: "amd64",
    arm: "arm",
  },
  platform: {
    darwin: "darwin",
    linux: "linux",
    win32: "windows",
  },
};

const binName = "findy-agent-cli";
const targetPath = "./bin";

const exec = (cmd) => {
  const exec = require("child_process").exec;
  return new Promise((resolve, reject) => {
    exec(cmd, (error, stdout, stderr) => {
      if (error) {
        console.warn(error);
      }
      resolve(stdout ? stdout : stderr);
    });
  });
};

const run = async (targetPath) => {
  const { platform, arch } = process;
  if (!platform in supported.platform || !arch in supported.arch) {
    console.error(`Not supported: ${platform}, ${arch}`);
    process.exit(1);
  }
  const path = `./dist/${binName}_${supported.platform[platform]}_${supported.arch[arch]}/${binName}`;
  await exec(`cp ${path} ${targetPath}/${binName}`);
};

if (!fs.existsSync(targetPath)) {
  fs.mkdirSync(targetPath);
}

run(targetPath);
