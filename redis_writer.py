import csv
import redis


def read_csv(file_path):
    redis_client = redis.StrictRedis(host='127.0.0.1', port=6379, db=0)
    main_dict = {}
    with open(file_path, newline='', encoding='utf-8') as csvfile:
        reader = csv.reader(csvfile)
        for row in reader:
            inner_dict = {}
            inner_dict["city"] = row[0]
            inner_dict["city_ascii"] = row[1]
            inner_dict["lat"] = row[2]
            inner_dict["lng"] = row[3]
            inner_dict["country"] = row[4]
            inner_dict["iso2"] = row[5]
            inner_dict["iso3"] = row[6]
            inner_dict["admin_name"] = row[7]
            inner_dict["capital"] = row[8]
            inner_dict["population"] = row[9]
            inner_dict["id"] = row[10]

            city = inner_dict["city_ascii"].replace(" ", "_").lower()

            if city in main_dict:
                saved_data = main_dict[city]
                if saved_data["population"] == "":
                    saved_data["population"] = "0"

                if inner_dict["population"] == "":
                    inner_dict["population"] = "0"

                saved_population = float(saved_data["population"])
                current_population = float(inner_dict["population"])

                if current_population > saved_population:
                    main_dict[city] = inner_dict
            
            else:
                main_dict[city] = inner_dict
            
    for key, value in main_dict.items():
        keyname = key.lower() + "_simple_street_maps"
        redis_client.hmset(keyname, value)
        print("Setting:Key->", key, " Values->", value)


read_csv("worldcities.csv")