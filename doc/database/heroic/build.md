# Build

- clone and cd into project dir

````bash
git clone git@github.com:spotify/heroic.git
cd heroic
# they shaded some dependencies like bigtable
tools/install-repackaged
mvn package
````