FROM codercom/code-server:3.0.2

USER root
ARG GO_VERSION=1.14.1
# Install go and kaniko to build images
RUN apt-get update \
    && apt install wget \
    && wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -xzf go${GO_VERSION}.linux-amd64.tar.gz -C /usr/local \
    && apt-get install gcc -y \
    && rm -r go${GO_VERSION}.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

# Install programming tools
RUN apt-get install -y python3-pip python3-venv \
    && update-alternatives --install /usr/bin/python python /usr/bin/python3 1

ENV PATH=$PATH:/home/coder/.local/bin

# Install kaniko-build tool

COPY kaniko_build /home/coder/kaniko_build
RUN pip3 install /home/coder/kaniko_build \
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

RUN pip3 install pyspark==2.4.5 \
    findspark \
    pyarrow==0.14.1 \
    mlflow==1.2.0 \
    streamlit

RUN apt-get install -y nodejs npm
RUN pip3 install jupyter jupyterlab

RUN chown -R coder:coder /usr/local/share/jupyter/

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