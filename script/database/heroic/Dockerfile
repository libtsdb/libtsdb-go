FROM maven:3.5.2-jdk-8

RUN git clone https://github.com/spotify/heroic.git \
    && cd heroic \
    && tools/install-repackaged \
    && mvn package

# TODO: $ java -cp $PWD/heroic-dist/target/heroic-dist-0.0.1-SNAPSHOT-shaded.jar com.spotify.heroic.HeroicService <config>