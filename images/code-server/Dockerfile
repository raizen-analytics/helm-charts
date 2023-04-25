FROM codercom/code-server

USER root
ARG GO_VERSION=1.14.1
ARG PYTHON_VERSION=3.10.0
# Install cron
RUN apt-get update && apt-get install -y cron
COPY cronjob /etc/cron.d/cronjob
RUN chmod 0644 /etc/cron.d/cronjob && crontab /etc/cron.d/cronjob
# Install go and kaniko to build images
RUN apt-get update \
    && apt install wget \
    && wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -xzf go${GO_VERSION}.linux-amd64.tar.gz -C /usr/local \
    && apt-get install gcc -y \
    && rm -r go${GO_VERSION}.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

# Install programming tools
RUN apt-get install wget build-essential libreadline-dev libncursesw5-dev libssl-dev libsqlite3-dev tk-dev libgdbm-dev libc6-dev libbz2-dev libffi-dev zlib1g-dev -y
RUN apt-get install build-essential. build-essential -y
RUN wget https://www.python.org/ftp/python/${PYTHON_VERSION}/Python-${PYTHON_VERSION}.tgz

RUN tar -xvf Python-${PYTHON_VERSION}.tgz

WORKDIR ./Python-${PYTHON_VERSION}

RUN ./configure --enable-optimizations

RUN make altinstall
RUN apt-get clean \
    && apt-get autoremove --purge \
    && apt-get remove python3.9 -y \
    && apt-get autoremove -y
RUN which python3.10
RUN update-alternatives --install /usr/bin/python python /usr/local/bin/python3.10  1 \
    && update-alternatives --install /usr/bin/pip pip /usr/local/bin/pip3.10 1
RUN python -v
ENV PATH=$PATH:/home/coder/.local/bin

#Install ZSH
RUN apt-get update
RUN apt-get install -y zsh
RUN wget https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh -O - | zsh || true

# Install kaniko-build tool

COPY kaniko_build /home/coder/kaniko_build
RUN pip install /home/coder/kaniko_build \
    && rm -r /home/coder/kaniko_build

# Install Spark and python tools

COPY --from=java:8 /usr/lib/jvm/java-8-openjdk-amd64 /usr/lib/jvm/java-8-openjdk-amd64
ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-amd64
RUN apt-get install -y ca-certificates-java

ARG HADOOP_VERSION="3.2.1"
ENV HADOOP_HOME "/opt/hadoop"
RUN curl https://archive.apache.org/dist/hadoop/core/hadoop-${HADOOP_VERSION}/hadoop-${HADOOP_VERSION}.tar.gz \
    | tar xz -C /opt && mv /opt/hadoop-${HADOOP_VERSION} ${HADOOP_HOME}
ENV HADOOP_COMMON_HOME "${HADOOP_HOME}"
ENV HADOOP_CLASSPATH "${HADOOP_HOME}/share/hadoop/tools/lib/*"
ENV HADOOP_CONF_DIR "${HADOOP_HOME}/etc/hadoop"
ENV PATH "$PATH:${HADOOP_HOME}/bin"
ENV HADOOP_OPTS "$HADOOP_OPTS -Djava.library.path=${HADOOP_HOME}/lib"
ENV HADOOP_COMMON_LIB_NATIVE_DIR "${HADOOP_HOME}/lib/native"
ENV YARN_CONF_DIR "${HADOOP_HOME}/etc/hadoop"
# install Spark
ARG SPARK_VERSION="2.4.5"
ARG PY4J_VERSION="0.10.7"
ENV SPARK_HOME "/opt/spark"
RUN curl https://archive.apache.org/dist/spark/spark-${SPARK_VERSION}/spark-${SPARK_VERSION}-bin-without-hadoop.tgz \
    | tar xz -C /opt && mv /opt/spark-${SPARK_VERSION}-bin-without-hadoop ${SPARK_HOME}
ENV PATH "$PATH:${SPARK_HOME}/bin"
ENV LD_LIBRARY_PATH "${HADOOP_HOME}/lib/native"
ENV SPARK_DIST_CLASSPATH "${HADOOP_HOME}/etc/hadoop\
:${HADOOP_HOME}/share/hadoop/common/lib/*\
:${HADOOP_HOME}/share/hadoop/common/*\
:${HADOOP_HOME}/share/hadoop/hdfs\
:${HADOOP_HOME}/share/hadoop/hdfs/lib/*\
:${HADOOP_HOME}/share/hadoop/hdfs/*\
:${HADOOP_HOME}/share/hadoop/yarn/lib/*\
:${HADOOP_HOME}/share/hadoop/yarn/*\
:${HADOOP_HOME}/share/hadoop/mapreduce/lib/*\
:${HADOOP_HOME}/share/hadoop/mapreduce/*\
:${HADOOP_HOME}/share/hadoop/tools/lib/*\
:${HADOOP_HOME}/contrib/capacity-scheduler/*.jar"
ENV PYSPARK_PYTHON "/usr/bin/python"
ENV PYTHONPATH "${SPARK_HOME}/python:${SPARK_HOME}/python/lib/py4j-${PY4J_VERSION}-src.zip:${PYTHONPATH}"
ENV SPARK_OPTS "--driver-java-options=-Xms1024M --driver-java-options=-Xmx4096M --driver-java-options=-Dlog4j.logLevel=info"

RUN pip install pyspark==2.4.5 \
    findspark \
    pyarrow==0.14.1 \
    mlflow==1.2.0 \
    streamlit \
    apache-airflow==1.10.10 \
    apache-airflow[kubernetes]==1.10.10 \
    apache-airflow[ssh]==1.10.10 \
    psycopg2-binary \
    azure-datalake-store \
    pymemcache \
    lazy-import==0.2.2

RUN apt-get install -y nodejs npm openssh-server
RUN pip install jupyter jupyterlab

# install data data science tools

RUN pip install wheel \
    mlflow==1.7.2 \
    pandas==1.0.3 \
    numpy==1.18.2 \
    deap==1.3.1 \
    tensorflow==2.1.0 \
    tensorflow-cpu==2.1.0 \
    keras==2.3.1 \
    joblib==0.14.1 \
    matplotlib==3.2.1

RUN chown -R coder:coder /usr/local/share/jupyter/

#Install kubectl

RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl \
    && chmod +x kubectl \
    && mv kubectl /usr/local/bin/kubectl

#Install R-Studio
# install prerequisities of r-studio server
# install r-studio server for derbian 10

USER root
RUN apt-get install -y r-base gdebi-core lib32gcc1 lib32stdc++6 \
    libclang-7-dev libclang-common-7-dev libclang-dev libclang1-7 \
    libgc1c2 libobjc-8-dev libobjc4 libpq5 \
    && wget https://download2.rstudio.org/server/bionic/amd64/rstudio-server-1.4.1106-amd64.deb \
    && gdebi -n rstudio-server-1.4.1106-amd64.deb \
    && rm rstudio-server-1.4.1106-amd64.deb

RUN echo 'directory=/home/coder/rstudiodb' >> /etc/rstudio/database.conf \
    && chmod -R +r /etc/rstudio/

ENV USER=coder
# COPY requirements.txt requirements.txt
# RUN pip3 install -r requirements.txt \
#     && rm requirements.txt

# lib mssql

# COPY ./libs/mssql-jdbc-7.2.2.jre8.jar ./libs/mssql-jdbc-7.2.2.jre8.jar
# RUN chown code:code ./libs/mssql-jdbc-7.2.2.jre8.jar

# Add custom entrypoint to change username
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh
USER coder

ENTRYPOINT [ "/usr/local/bin/entrypoint.sh" ]