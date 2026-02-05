{
  description = "Suger Go SDK";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.11";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        devShells.default = pkgs.mkShell {
          nativeBuildInputs = [ pkgs.bashInteractive ];
          hardeningDisable = [ "fortify" ];
          buildInputs = with pkgs;
            [
              # go programming tools
              go
              gotools # goimports godoc and others
              gopls
              delve

              # lint tools
              golangci-lint

              # build tools
              gnumake
              gh
            ];
          shellHook = ''
            if [[ "$(uname)" == "Darwin" && -d "/Library/Developer/CommandLineTools" ]]; then
              export DEVELOPER_DIR="/Library/Developer/CommandLineTools"
              export SDKROOT="/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk"
              export PATH="/Library/Developer/CommandLineTools/usr/bin:$PATH"
            fi

            # Add GOPATH/bin to PATH for tools installed via go install
            export PATH="$PATH:$(go env GOPATH)/bin"

            echo "Suger Go SDK development environment"
            echo "Go version: $(go version)"
          '';
        };
      });
}
