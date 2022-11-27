import random
import keyboard
from time import sleep
from multiprocessing.pool import ThreadPool


# task executed in a worker thread
def post(list,quit,number):
    print(f'post {number} started\n')
    while True:
        if keyboard.is_pressed("q"):
            print("You pressed q")
            break
        if len(quit)>0:
            break
        if len(list)>=100:
            continue
        if len(list)<=80:
            t=random.randint(1,100)
            print(f'post {number} gave {t}\n')
            list.append(t)

def consumer(list,quit,number):
    print(f'consumer {number} started\n')
    while True:
        if keyboard.is_pressed("q"):
            print("You pressed q")
            break
        if len(list) == 0 and len(quit)==0:
            continue
        if len(list) == 0 and len(quit)>0:
            break
        if len(list) > 1:
            print(f'consumer {number} get {list[-1]}\n')
            list.pop()


# initialize a worker in the thread pool
def initialize_worker():
    # report a message
    print('Initializing worker...')


# protect the entry point
if __name__ == '__main__':
    ban = []
    qu = []
    flag=0
    # create and configure the thread pool
    with ThreadPool(5, initializer=initialize_worker) as pool:
        # issue tasks to the thread pool
        for _ in range(3):
            if _==2:
                _ = pool.apply_async(post, args=(ban, qu, flag + 1))
            _ = pool.apply_async(post,args=(ban,qu,flag+1))
            _ = pool.apply_async(consumer,args=(ban,qu,flag+1))
            flag+=1
        # close the thread pool
        pool.close()
        # wait for all tasks to complete
        pool.join()
