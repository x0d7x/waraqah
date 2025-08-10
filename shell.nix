{
  pkgs ? import <nixpkgs> { },
}:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.fish
  ];

  shellHook = ''
    export SHELL=${pkgs.fish}/bin/fish
  '';
}
