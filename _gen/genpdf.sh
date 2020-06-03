echo 'Kill all Jekyll instances'
kill -9 $(ps aux | grep '[j]ekyll' | awk '{print $2}')
clear

cd "$(dirname "$0")"
cd ..

echo "Building PDF-friendly HTML site ...";
bundle exec jekyll serve --detach --config _config.yml,pdfconfigs/all_pdf.yml;
echo "done";

echo "Building the PDF ...";
prince --javascript --input-list=_site/pdfconfigs/prince-list.txt -o pdf/all.pdf;

echo "Done. Look in the pdf directory to see if it printed successfully."

