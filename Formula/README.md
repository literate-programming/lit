homebrew-lit
============
Homebrew formulae for installing lit cli tool on MacOS.

# Usage

Use the following commands to Tap the Homebrew repository

```sh
brew tap literate-programming/lit
```

Then install the `lit-cli` formula

```sh
brew install lit-cli
```

# Testing changes
The following process can be used to test changes made to formula:
1. Ensure the target formula has already been installed on your system. I will use `lit-cli` to demonstrate.
2. Fork this repo on Github and clone it locally to your machine.
3. Create a new branch for your changes.
4. Make your changes to the target formula and commit them.
5. Change directory into the `Formula` foldere
```sh
cd $(brew --repo literate-programming/lit)/Formula
```
6. Unlink the currently installed version of the formula with the following command:
```sh
brew unlink lit-cli
```
6. Copy your updated formula into the `literate-programming/lit/Formula` folder - replacing the exiting version
7. Test download of the new package from the Formula URL
```sh
brew fetch lit-cli
```
8. Test the installation using the new Formula (with verbose mode and based on the locally updated file)
```sh
brew install -vsd --git lit-cli
```

All going well, you should have just installed the updated version based on the new formula. If that is the case, submit a PR with your changes for further review. If you encounter issue, repeat the process above until you are able to successfully install the based on the udpated version.
