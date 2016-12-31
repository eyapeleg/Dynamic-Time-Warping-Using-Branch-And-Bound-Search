Dynamic Time Warping using Branch and Bound Search
---------------------------------------------------
This project aims to supply a different approach to the time warping measurement calculation. The time warping measures the distance between two time series in way that is more elastic so that different time points in both series could be associated with each other. This prevents from very similar time series events to be supposedly seen as very far away from eahc other when they have  a different starting point. Take sin(x) and cos(x). when comparing the distance on a regular scale they would be seen as very different time series instances although they have the same trends just with a different starting point. Here is where Time Warping shines, it is able to distinguish this behaviour and to single the these series are actually the same.  

Now, the initial approach to perform the calculation was described in Sakoe, Hiroaki, and Seibi Chiba. "Dynamic programming algorithm optimization for spoken word recognition". This approach uses a dynamic programming solution in order to perform the calculation. Although its execution is very quick ( O(mn) where m,n are the sizes of the time series) it suffers from space complexity issues (has to maintain a matrix of size M*N in it) and it is also bound to be single threaded.  

In this project I offer a different approach for calculating the time warping distance – Branch and Bound Search. This approach was selected due the ability to multi-thread the search and the low space complexity it requires in comaprison to the first apporach.

You will find here 3 sections:
1. The source code
2. An Analysis of the results using ipython notebook
3. An Article with the description of the experiment, results and conclusions.
