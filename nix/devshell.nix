{
  perSystem,
  pkgs,
  ...
}:
perSystem.self.godoc.overrideAttrs (old: {
  GOROOT = "${old.go}/share/go";
  nativeBuildInputs =
    old.nativeBuildInputs
    ++ [
      perSystem.gomod2nix.default
      pkgs.enumer
      pkgs.delve
      pkgs.pprof
      pkgs.gotools
      pkgs.golangci-lint
      pkgs.cobra-cli
    ];
  shellHook = ''
    # this is only needed for hermetic builds
    unset GO_NO_VENDOR_CHECKS GOSUMDB GOPROXY GOFLAGS
  '';
})
