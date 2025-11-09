FROM python:3.11-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY server.py .
# criamos data/ dentro do container no start
VOLUME ["/app/data"]
EXPOSE 5555
CMD ["python", "server.py"]
