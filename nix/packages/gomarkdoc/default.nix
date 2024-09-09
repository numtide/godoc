{
  flake,
  # We need the following pragma to ensure deadnix doesn't remove inputs.
  # This package is being called with newScope/callPackage, which means it is only being passed args it defines.
  # We do not use inputs directly in this file, but need it for passing to the tests.
  # deadnix: skip
  inputs,
  perSystem,
  system,
  pkgs,
  pname,
  ...
} @ args: let
  inherit (pkgs) go lib;
  fs = lib.fileset;
in
  perSystem.gomod2nix.buildGoApplication rec {
    inherit pname;
    # there's no good way of tying in the version to a git tag or branch
    # so for simplicity's sake we set the version as the commit revision hash
    # we remove the `-dirty` suffix to avoid a lot of unnecessary rebuilds in local dev
    version = lib.removeSuffix "-dirty" (flake.shortRev or flake.dirtyShortRev);

    # ensure we are using the same version of go to build with
    inherit go;

    src = fs.toSource {
      root = ../../..;
      fileset = fs.unions [
        ../../../cmd
        ../../../go.mod
        ../../../go.sum
        ../../../main.go
        ../../../pkg
      ];
    };

    modules = ./gomod2nix.toml;

    ldflags = [
      "-s"
      "-w"
      "-X github.com/numtide/gomarkdoc/pkg/build.Name=${pname}"
      "-X github.com/numtide/gomarkdoc/pkg/build.Version=v${version}"
    ];

    meta = with lib; {
      description = "gomarkdoc: Go Doc generation in Markdown format";
      homepage = "https://github.com/numtide/gomarkdoc";
      license = licenses.mit;
      mainProgram = "gomarkdoc";
    };
  }
