{
  description = "FlyDistSys";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, utils }:
    utils.lib.eachDefaultSystem (system:
      let
        pname = "fly-dist-sys";
        version = "0.0.1";
        pkgs = import nixpkgs {
          inherit system;
        };
        tools = with pkgs; [
          # https://github.com/golang/vscode-go/blob/master/docs/tools.md
          delve
          go-outline
          golangci-lint
          gomodifytags
          gopls
          gopkgs
          gotests
          impl
        ];

        jepsen-maelstrom = pkgs.stdenv.mkDerivation {
            name = "maelstrom";
            version = "0.2.3";
            src = builtins.fetchurl {
              url = "https://github.com/jepsen-io/maelstrom/releases/download/v0.2.3/maelstrom.tar.bz2";
              sha256 = "sha256:06jnr113nbnyl9yjrgxmxdjq6jifsjdjrwg0ymrx5pxmcsmbc911";
          };
          installPhase = ''
            mkdir -p $out/bin
            cp -r lib $out/bin/lib
            install -m755 -D maelstrom $out/bin/maelstrom
          '';
        };

        maelstrom-packages = with pkgs; [
          jepsen-maelstrom
          openjdk19
          graphviz
          gnuplot
        ];
      in
      rec {
        # `nix develop`
        devShell = with pkgs; mkShell {
          buildInputs = [ go gnumake ] ++ maelstrom-packages ++ tools;
        };
      });
}
