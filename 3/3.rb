class ThreadPool
  def initialize(size)
    @size = size
    @jobs = Queue.new
    @pool = Array.new(@size) do |i|
      Thread.new do
        Thread.current[:id] = i
        catch(:exit) do
          loop do
            job, args = @jobs.pop
            job.call(*args)
          end
        end
      end
    end
  end

  # add a job to queue
  def schedule(*args, &block)
    @jobs << [block, args]
  end

  # run threads and perform jobs from queue
  def run!
    @size.times do
      schedule { throw :exit }
    end
    @pool.map(&:join)
  end
end

def post(list,name)
  puts "post #{name} started"
  loop do
    t = rcislo
    if list.length >=100
      next
    end
    if list.length<=80
      list.push(t)
      puts "fabric gave #{list[-1]}"
    end
  end
end

def rcislo()
  rc = Random.new
  return rc.rand(100)
end

def consumer(list,name)
  puts "pocup #{name} started"
  loop do
    if list.length >0
      puts "store get #{list[-1]}"
      list.pop
    end
    if list.length==0
      next
    end
  end
end

# an instance of ThreadPool with 5 threads
pool = ThreadPool.new(5)
b =[]


2.times do |r|
  pool.schedule do
    sleep_time = rand(4) + 2
    sleep(sleep_time)
    consumer(b,r)
  end
end

# add 3 tasks to query
3.times do |i|
  pool.schedule do
    sleep_time = rand(4) + 2
    sleep(sleep_time)
    post(b,i)
    if i==2
      post(b,i)
    end
  end
end

# run all threads
pool.run!
