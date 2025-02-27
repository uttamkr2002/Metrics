# Metrics
Calculate the Metrics and Store in DB


# client

1> Fetch the Data from OS using gopsUtil Package
2> Store the data in DB

# Create Table in PostGreSQl

CREATE TABLE metrics_data (
    id SERIAL PRIMARY KEY,
    cpu_usage FLOAT NOT NULL,
    disk_percent FLOAT NOT NULL,
    used_disk BIGINT NOT NULL,
    free_disk BIGINT NOT NULL,
    total_disk BIGINT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


# Data Stored In PostGrESQL
![image](https://github.com/user-attachments/assets/f047159f-4d54-40ae-a5cd-dd3243055508)


# Added this feature for Visualization
![image](https://github.com/user-attachments/assets/6f71da53-c583-4ca3-b774-4c2680d541a1)
