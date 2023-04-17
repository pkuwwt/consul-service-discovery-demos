FROM python:3.10-slim
WORKDIR /dist

EXPOSE 5000
ADD requirements.txt /dist/
RUN pip install -r requirements.txt

ADD app.py /dist/

ENTRYPOINT ["python3", "-m", "app"]
