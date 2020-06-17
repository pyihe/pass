package main

const (
	commandGen = "gen" //新增
	commandGet = "get" //查询
	commandDel = "del" //删除
	commandSet = "set" //刷新

	help = `	

	command:
		gen	generate one new password for specified key,
			the key must not be empty. you can choose to 
			reset the password if the key already exist.

		get	get password of the specified key, all passwords
			will be list if key is empty.

		del	delete password of specified key, key must not
			be empty.

		set	reset the password of the specified key with your 
			own password. it will generate new password if 
			the given password is empty.

	key:		the key for password.

	pass:		the password you want to reset to the key.


`
)

var (
	fileName string
	code     = [62]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}
	chars = [20]byte{'~', '!', '@', '#', '$', '%', '^', '&', '*', '_', '-', '+', '=', '<', '>', ',', '.', '|', '/', ';'}
)
