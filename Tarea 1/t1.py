"""
    By Luis Felipe Alvarez Sanchez
    A01194173
    2/18/2021
"""

class Stack:
  # Constructor
  def __init__(self):
      self.stack = []
      self.topValue = -1
  # Inserts an element into the stack
  def push(self, val):
    self.topValue = val
    print(val, 'was pushed into the stack.')
    self.stack.append(val)
 # Checks if the stack is empty
  def isEmpty(self):
      return len(self.stack) <= 0
 # Prints the peek of the stack 
  def top(self):
      if not(len(self.stack) <= 0):
        print(self.topValue)
      else:
        print('Stack is empty')  
  # Pops an element from the stack
  def pop(self):
      if not(len(self.stack) <= 0):
         print(self.stack[len(self.stack) - 1], 'was removed from the stack.')
         self.stack = self.stack[:len(self.stack) -1] 
      else:
         print('Stack is empty')  
  # Prints the current stack
  def print_stack(self):
    print('Current stack:', self.stack)


class Queue:
  # Constructor
  def __init__(self):
      self.queue = []
      self.frontValue = -1
  # Inserts an element into the queue
  def push(self, val):
    print(val, 'was pushed into the queue.')
    self.queue.append(val)
    self.frontValue = self.queue[0]
 # Checks if the queue is empty
  def isEmpty(self):
      return len(self.queue) <= 0
 # Prints the front of the queue 
  def front(self):
      if not(len(self.queue) <= 0):
        print(self.queue[0])
      else:
        print('Queue is empty')  
  # Pops an element from the queue
  def pop(self):
      if not(len(self.queue) <= 0):
         print(self.queue[0], 'was removed from the queue.')
         self.queue = self.queue[1:] 
      else:
         print('Queue is empty')  
  # Prints the current queue
  def print_queue(self):
    print('Current queue:', self.queue)


# Stack Implementation
p1 = Stack()
p1.top()
p1.push(2)
p1.push(3)
p1.push(4)
p1.print_stack()
p1.pop()
p1.print_stack()

# Queue implementation
p2 = Queue()
p2.front()
p2.pop()
p2.print_queue()
p2.push(1)
p2.push(2)
p2.push(3)
p2.front()
p2.pop()
p2.front()
p2.print_queue()

# DICTIONARY IN PYTHON
# Methods ---------------------------------------------------------------
# clear -> removes the elements of the dictionary
# copy -> returns a copy of the dictionary
# fromkeys ->returns a dictionary with the specified key and value
# get -> returns the value of the specified key
# items -> returns a list containing a tuple for each key value pair
# keys -> returns a list containing the dictionary's keys
# pop -> removes the element with the specified key
# popitem -> removes the last inserted key value pair
# update -> updates the dictionary with the specified key value pairs
# values -> returns a list of all the values in the dictionary   

# Declaration of a Dictionary
dictionary_empty = {}
dictionary = {
    "name": "Luis",
    "age" : 22,
    "favorite_burgers":["Carls Jr", "In N Out" , "Five guys"]
}

# Display a dictionary
print(dictionary)

# Access specific value of dictionary
print(dictionary['name'])
print(dictionary['favorite_burgers'])

# Access specific value of dictionary and access one element.
print(dictionary['favorite_burgers'][0])

# Duplicates are not allowed in dictionaries
dictionary_example = {
   "name": "Luis",
   "name": "Felipe",
    "age" : 22,
    "favorite_burgers":["Carls Jr", "In N Out" , "Five guys"]
}
print(dictionary_example)

# Get the length of dictionary
print(len(dictionary))

# Get value of element in dictionary
x = dictionary['name']
y = dictionary.get('name')

# Get keys from dictionary
keys = dictionary.keys()
print(keys)

# Get all values from dictionary
values = dictionary.values()
print(values)

# Get the key value pairs of the dictionary
items = dictionary.items()
print(items)

# Update Dictionary
dictionary['name'] = 'Felipe'
print(dictionary['name'])
dictionary.update({"name": "Luis"})
print(dictionary['name'])

# Adding items into the dictionary
dictionary['sex'] = 'Male'
print(dictionary['sex'])
dictionary.update({"sex": "Male"})
print(dictionary['sex'])

# Removing elements of the dictionary
dictionary.pop('sex')
print(dictionary)

# Removes the last inserted element
dictionary.popitem()
print(dictionary)

# Removes name of the dictionary
del dictionary['name']

#Copies the dictionary into dictionary_copy
dictionary_copy = dictionary.copy()

# Clears the dictionary
dictionary_copy.clear()
print(dictionary_copy)

# Print dictionary key names
for x in dictionary:
    print(x)

# Print dictionary values
for x in dictionary:
    print(dictionary[x])

# Reference: https://www.w3schools.com/python/python_dictionaries_methods.asp