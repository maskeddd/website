{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      nixpkgs,
      flake-utils,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gopls
            tailwindcss_4

            (writeShellScriptBin "templ" ''
              exec go tool templ "$@"
            '')
            (writeShellScriptBin "air" ''
              exec go tool air "$@"
            '')
          ];
        };
      }
    );
}
