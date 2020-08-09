import mariadb

class Samaria:
	def __init__(self, password, database, host='0.0.0.0', user='root'):
		self.conn	= mariadb.connect(host=host, user=user, password=password, database=database)
		self.cursor = self.conn.cursor()
	
	# async to be implemented, there is no advantage on Samaria over sync Mariadb at the moment
	async def execute(self, query):
		self.cursor.execute(query)

	async def fetchall(self):
		return self.cursor.fetchall()
	
	async def commit(self):
		self.conn.commit()

if __name__ == "__main__":
	samaria = Samaria(password='worbdj12', database='devqueue')
	print(samaria)
